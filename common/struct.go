package common

import (
	"strings"
)

type Event struct {
	Departement Departement
	City        string
	StationID   string

	VoteID   string
	VoteName string

	Register   uint
	Abstention uint
	Blank      uint
	Null       uint

	Option []Option
}

// A polling zone, all France, a department, a city or a stationID.
type Zone struct {
	Departement Departement
	City        string
	StationID   string

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

func ConstOptions(args ...string) (options []Option) {
	options = make([]Option, 0, len(args))
	for i, a := range args {
		s := strings.Split(a, "\t")
		opinion := Opinion(0)
		switch s[0] {
		case "XL":
			opinion = OpinionFarLeft
		case "L":
			opinion = OpinionLeft
		case "C":
			opinion = OpinionCenter
		case "R":
			opinion = OpinionRight
		case "XR":
			opinion = OpinionFarRight
		case "O":
			opinion = OpinionOther
		}
		gender := GenderMan
		switch s[2] {
		case "M":
		case "F", "W":
			gender = GenderWoman
		default:
			panic("unknow gender: " + s[2])
		}
		options = append(options, Option{
			Position: uint(i) + 1,
			Party:    s[1],
			Opinion:  opinion,
			Name:     s[3],
			Gender:   gender,
		})
	}
	return
}

type SubSummary struct {
	Group string
	Summary
}

type Summary struct {
	Register uint
	Result   [OpinionLength]uint
}
