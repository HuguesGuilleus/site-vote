package votation

import (
	"cmp"
	"slices"
	"time"
)

// Polling station (fr: bureau de vote)
type Station struct {
	Departement Departement
	City        string
	CodeStation string
	Votation    []Votation
}

type Votation struct {
	Name string
	Date time.Time
	VotationResult
}

type VotationResult struct {
	Register   uint
	Abstention uint
	Blank      uint
	Null       uint
	Result     []Result
}

type Result struct {
	*Option
	Result uint
}

type Option struct {
	Position uint
	Party    string
	Opinion  Opinion
	Name     string
	Gender   Gender
}

type Summary struct {
	Register uint
	Result   [OpinionLength]uint
}

func NewSummary(v *VotationResult) (s Summary) {
	s.Register = v.Register
	s.Result[OpinionAbstention] = v.Abstention
	s.Result[OpinionBlank] = v.Blank
	s.Result[OpinionNull] = v.Null
	for _, r := range v.Result {
		s.Result[r.Opinion] += r.Result
	}
	return
}

func MergeStation(stationsArgs ...[]*Station) []*Station {
	allStations := make([]*Station, 0)
	for _, stations := range stationsArgs {
		allStations = append(allStations, stations...)
	}

	slices.SortFunc(allStations, func(a, b *Station) int {
		return cmp.Or(
			cmp.Compare(a.Departement, b.Departement),
			cmp.Compare(a.City, b.City),
			cmp.Compare(a.CodeStation, b.CodeStation),
		)
	})

	if len(allStations) <= 1 {
		return allStations
	}

	w := 0
	for _, s := range allStations[1:] {
		p := allStations[w]
		if p.Departement == s.Departement && p.City == s.City && p.CodeStation == s.CodeStation {
			p.Votation = append(p.Votation, s.Votation...)
		} else {
			w++
			allStations[w] = s
		}
	}
	w++
	clear(allStations[w:])
	return allStations[:w]
}
