package common

import (
	"cmp"
	"iter"
	"slices"
	"sniffle/tool"
	"sync"
)

func Call(t *tool.Tool, votations ...func(t *tool.Tool) []*Event) []*Event {
	all := make([]*Event, 0)

	allMutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(votations))
	for _, v := range votations {
		go func() {
			defer wg.Done()
			events := v(t)
			for _, e := range events {
				slices.SortFunc(e.Option, optionCompare)
			}
			allMutex.Lock()
			defer allMutex.Unlock()
			all = append(all, events...)
		}()
	}
	wg.Wait()

	slices.SortFunc(all, func(a, b *Event) int {
		return cmp.Or(
			cmp.Compare(a.Departement, b.Departement),
			cmp.Compare(a.City, b.City),
			cmp.Compare(len(a.StationID), len(b.StationID)),
			cmp.Compare(a.StationID, b.StationID),
		)
	})

	return all
}

func AllFrance(events []*Event) *Zone {
	return &Zone{
		Vote: mergeVote(events, nil),
		Sub:  mergeSub(events, func(e *Event) string { return e.Departement.Code() }),
	}
}

// A filter function to check if the future zone is necessary.
type Skip = func(Departement, string) bool

func ByDepartement(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		zoneEvents := []*Event(nil)
		for len(events) != 0 {
			zoneEvents, events = splitEvent(events, func(a, b *Event) bool {
				return a.Departement != b.Departement
			})
			e0 := zoneEvents[0]
			if skip(e0.Departement, "") {
				continue
			}
			getGroup := func(e *Event) string { return e.City }
			z := &Zone{
				Departement: e0.Departement,
				City:        e0.City,
				Vote:        mergeVote(zoneEvents, nil),
				Sub:         mergeSub(zoneEvents, getGroup),
			}
			if !yield(z) {
				return
			}
		}
	}
}

func ByCity(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		zoneEvents := []*Event(nil)
		for len(events) != 0 {
			zoneEvents, events = splitEvent(events, func(a, b *Event) bool {
				return a.Departement != b.Departement || a.City != b.City
			})
			e0 := zoneEvents[0]
			if skip(e0.Departement, e0.City) {
				continue
			}
			getGroup := func(e *Event) string { return e.StationID }
			z := &Zone{
				Departement: e0.Departement,
				City:        e0.City,
				Vote:        mergeVote(zoneEvents, getGroup),
				Sub:         mergeSub(zoneEvents, getGroup),
			}
			if !yield(z) {
				return
			}
		}
	}
}
func ByStation(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		zoneEvents := []*Event(nil)
		for len(events) != 0 {
			zoneEvents, events = splitEvent(events, func(a, b *Event) bool {
				return a.Departement != b.Departement || a.City != b.City || a.StationID != b.StationID
			})
			e0 := zoneEvents[0]
			if skip(e0.Departement, e0.City) {
				continue
			}
			z := &Zone{
				Departement: e0.Departement,
				City:        e0.City,
				StationID:   e0.StationID,
				Vote:        mergeVote(zoneEvents, nil),
			}
			if !yield(z) {
				return
			}
		}
	}
}

func splitEvent(
	events []*Event,
	notSame func(a, b *Event) bool) ([]*Event, []*Event) {
	if len(events) == 0 {
		return events, nil
	}

	for i, e := range events[1:] {
		if notSame(events[0], e) {
			return events[:i+1], events[i+1:]
		}
	}

	return events, nil
}

func mergeVote(events []*Event, getGroup func(*Event) string) (votes []Vote) {
	mvotes := make(map[string]*Vote)
	for _, e := range events {
		v := mvotes[e.VoteID]
		if v == nil {
			v = &Vote{ID: e.VoteID, Name: e.VoteName, Option: make([]Option, 0)}
			mvotes[e.VoteID] = v
		}

		v.Register += e.Register
		v.Abstention += e.Abstention
		v.Blank += e.Blank
		v.Null += e.Null
		v.Option = mergeOption(v.Option, e.Option)

		sum := e.Sum()
		v.Summary.Add(sum)

		if getGroup != nil {
			group := getGroup(e)
			if i := len(v.SubSummary) - 1; i == -1 || v.SubSummary[i].Group != group {
				v.SubSummary = append(v.SubSummary, SubSummary{group, sum})
			} else {
				v.SubSummary[i].Summary.Add(sum)
			}
		}
	}

	votes = make([]Vote, 0, len(mvotes))
	for _, v := range mvotes {
		votes = append(votes, *v)
	}
	slices.SortFunc(votes, func(a, b Vote) int { return cmp.Compare(b.ID, a.ID) })
	return
}

func mergeSub(events []*Event, getGroup func(*Event) string) (sub []string) {
	sub = make([]string, 0)
	for _, e := range events {
		if g := getGroup(e); len(sub) == 0 || sub[len(sub)-1] != g {
			sub = append(sub, g)
		}
	}
	return
}

// Merge options inside out.
func mergeOption(out, add []Option) []Option {
	if out == nil || len(out)+len(add) > 50 {
		return nil
	}

	for w := 0; len(add) != 0 && w < len(out); w++ {
		switch optionCompare(out[w], add[0]) {
		case 0:
			out[w].Result += add[0].Result
			add = add[1:]
		case 1:
			out = slices.Insert(out, w, add[0])
			add = add[1:]
		case -1: // do nothing
		}
	}

	return append(out, add...)
}

// Compare two option with their opinion name and position
//
//	-1: a < b
//	 0: a == b
//	+1: a > b
func optionCompare(a Option, b Option) int {
	return cmp.Or(
		cmp.Compare(a.Opinion, b.Opinion),
		cmp.Compare(a.Name, b.Name),
		cmp.Compare(a.Position, b.Position),
	)
}

func (e *Event) Sum() (sum Summary) {
	sum.Register = e.Register
	sum.Result[OpinionAbstention] = e.Abstention
	sum.Result[OpinionBlank] = e.Blank
	sum.Result[OpinionNull] = e.Null
	for _, o := range e.Option {
		sum.Result[o.Opinion] += o.Result
	}
	return
}

func (s *Summary) Add(t Summary) {
	s.Register += t.Register
	for i, r := range t.Result {
		s.Result[i] += r
	}
}
