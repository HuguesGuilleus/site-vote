package driver2024legislative

type Opinion byte

const (
	OpinionOther Opinion = iota
	OpinionFarLeft
	OpinionLeft
	OpinionCenter
	OpinionRight
	OpinionFarRight

	OpinionBlank
	OpinionNull
	OpinionAbstention

	OpinionLength
)

func (o Opinion) String() string {
	switch o {
	case OpinionOther:
		return "A"
	case OpinionFarLeft:
		return "EG"
	case OpinionLeft:
		return "G"
	case OpinionCenter:
		return "C"
	case OpinionRight:
		return "D"
	case OpinionFarRight:
		return "ED"

	case OpinionBlank:
		return "B"
	case OpinionNull:
		return "N"
	case OpinionAbstention:
		return "A"

	default:
		return "X"
	}
}
