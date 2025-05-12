package votation

type Gender byte

const (
	GenderList Gender = iota
	GenderMan
	GenderWoman
)

// Return french pronun in parenthesis or empty string if otthneer.
func (g Gender) String() string {
	switch g {
	case GenderMan:
		return "(il)"
	case GenderWoman:
		return "(elle)"
	default:
		return ""
	}
}
