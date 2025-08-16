package common

type Event struct {
	Departement Departement
	City        string
	StationID   string
	District    string

	VoteID   string
	VoteName string

	Register   uint
	Abstention uint
	Blank      uint
	Null       uint

	Option []Option

	// True if a you can vote to multuple option.
	// Example: Municipale élection in France where max 999 citizens live in the city.
	SplitVoting bool
}

// A polling zone, all France, a department, a city or a stationID.
type Zone struct {
	Departement Departement
	City        string
	StationID   string
	District    string

	Parents []string

	// Element with the same level.
	// Example: with Departement Zone, a legislative district.
	Same []string

	Sub []string

	Vote []Vote
}

type Vote struct {
	ID   string
	Name string

	Register   uint
	Abstention uint
	Blank      uint
	Null       uint
	Option     []Option

	Summary    Summary
	SubSummary []SubSummary
}

type Option struct {
	Result   uint
	Position uint
	Party    string
	Opinion  Opinion
	Name     string
	Gender   Gender
}

type SubSummary struct {
	Group string
	Summary
}

type Summary struct {
	Register uint
	Result   [OpinionLength]uint
}
