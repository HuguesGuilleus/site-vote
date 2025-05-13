package votation

import "time"

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
	Position  uint
	Party     string
	Opinion   Opinion
	FirstName string
	LastName  string
	Gender    Gender
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
