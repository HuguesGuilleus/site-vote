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
			cmp.Compare(b.VoteID, a.VoteID),
		)
	})

	return all
}

func AllFrance(events []*Event) *Zone {
	z := new(Zone)
	z.mergeOneZone(events, func(e *Event) string { return e.Departement.Code() })
	return z
}

// A filter function to check if the future zone is necessary.
type Skip = func(Departement, string) bool

func ByDepartement(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		by(yield, events,
			func(e *Event) string { return e.City },
			func(a, b *Event) bool {
				return a.Departement != b.Departement
			},
			func(e *Event) *Zone {
				if skip(e.Departement, "") {
					return nil
				}
				return &Zone{
					Departement: e.Departement,
				}
			},
		)
	}
}
func ByCity(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		by(yield, events,
			func(e *Event) string { return e.StationID },
			func(a, b *Event) bool {
				return a.Departement != b.Departement || a.City != b.City
			},
			func(e *Event) *Zone {
				if skip(e.Departement, e.City) {
					return nil
				}
				return &Zone{
					Departement: e.Departement,
					City:        e.City,
				}
			},
		)
	}
}
func ByStation(events []*Event, skip Skip) iter.Seq[*Zone] {
	return func(yield func(*Zone) bool) {
		by(yield, events,
			nil,
			func(a, b *Event) bool {
				return a.Departement != b.Departement ||
					a.City != b.City ||
					a.StationID != b.StationID
			},
			func(e *Event) *Zone {
				if skip(e.Departement, e.City) {
					return nil
				}
				return &Zone{
					Departement: e.Departement,
					City:        e.City,
					StationID:   e.StationID,
				}
			},
		)
	}
}

func by(
	yield func(*Zone) bool,
	events []*Event,
	getGroup func(*Event) string,
	notSame func(a, b *Event) bool,
	newZone func(*Event) *Zone,
) {
	switch len(events) {
	case 0:
		return
	case 1:
		if z := newZone(events[0]); z != nil {
			z.mergeOneZone(events, getGroup)
			yield(z)
		}
		return
	}

	for i, e := range events[1:] {
		if notSame(events[0], e) {
			if z := newZone(events[0]); z != nil {
				z.mergeOneZone(events[:i+1], getGroup)
				if !yield(z) {
					return
				}
			}
			by(yield, events[i+1:], getGroup, notSame, newZone)
			return
		}
	}

	if z := newZone(events[0]); z != nil {
		z.mergeOneZone(events, getGroup)
		yield(z)
	}
}

func (z *Zone) mergeOneZone(events []*Event, getGroup func(*Event) string) {
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

	if getGroup != nil {
		z.Sub = make([]string, 0)
		for _, e := range events {
			if g := getGroup(e); len(z.Sub) == 0 || z.Sub[len(z.Sub)-1] != g {
				z.Sub = append(z.Sub, g)
			}
		}
	}

	z.Vote = make([]Vote, 0, len(mvotes))
	for _, v := range mvotes {
		z.Vote = append(z.Vote, *v)
	}
	slices.SortFunc(z.Vote, func(a, b Vote) int {
		return cmp.Compare(b.ID, a.ID)
	})
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
