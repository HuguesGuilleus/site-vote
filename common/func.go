package common

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Merges all events, and sort it.
func Merge(args ...[]*Event) (events []*Event) {
	l := 0
	for _, a := range args {
		l += len(a)
	}
	events = make([]*Event, 0, l)
	for _, a := range args {
		events = append(events, a...)
	}

	slices.SortFunc(events, func(a, b *Event) int {
		return cmp.Or(
			cmp.Compare(a.Departement, b.Departement),
			cmp.Compare(a.City, b.City),
			cmp.Compare(a.StationID, b.StationID),
			cmp.Compare(b.VoteID, a.VoteID),
		)
	})

	for _, e := range events {
		slices.SortFunc(e.Option, optionCompare)
	}

	return
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
				z.mergeOneZone(events[:i+2], getGroup)
				if !yield(z) {
					return
				}
			}
			by(yield, events[i+2:], getGroup, notSame, newZone)
			return
		}
	}

	if z := newZone(events[0]); z != nil {
		z.mergeOneZone(events, getGroup)
		yield(z)
	}
}

func (z *Zone) mergeOneZone(events []*Event, getGroup func(*Event) string) {
	msub := make(map[string]struct{})
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
		if getGroup == nil {
			continue
		}
		group := getGroup(e)
		msub[group] = struct{}{}
		if i := len(v.SubSummary) - 1; i > 0 && v.SubSummary[i].Group == group {
			v.SubSummary[i].Summary.Add(sum)
		} else {
			v.SubSummary = append(v.SubSummary, SubSummary{group, sum})
		}
	}

	if len(msub) > 0 {
		z.Sub = slices.AppendSeq(make([]string, 0, len(msub)), maps.Keys(msub))
		slices.Sort(z.Sub)
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
	if out == nil || len(out)+len(add) > 200 {
		return nil
	}

	out = append(out, add...)
	if len(out) <= 1 {
		return out
	}
	slices.SortFunc(out, optionCompare)

	w := 0
	for _, o := range out[1:] {
		if optionCompare(out[w], o) == 0 {
			out[w].Result += o.Result
		} else {
			w++
			out[w] = o
		}
	}

	return out[:w+1]
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
