package common

import "strings"

// Parse option for a national votation with people (presidential).
// Format (use tabulation): OPINION PARTY GENDER NAME.
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
		gender := GenderList
		switch s[2] {
		case "L":
		case "M":
			gender = GenderMan
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

// Parse option for a constant national votation with list (like EU).
// Format (use tabulation): OPINION PARTY NAME.
func ConstOptionsList(args ...string) (options []Option) {
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
		options = append(options, Option{
			Position: uint(i) + 1,
			Party:    s[1],
			Opinion:  opinion,
			Name:     s[2],
			Gender:   GenderList,
		})
	}
	return
}
