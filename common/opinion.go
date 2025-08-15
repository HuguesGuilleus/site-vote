package common

type Opinion byte

const (
	OpinionFarLeft Opinion = iota
	OpinionLeft
	OpinionCenter
	OpinionRight
	OpinionFarRight
	OpinionOther

	OpinionBlank
	OpinionNull
	OpinionAbstention

	OpinionLength
)

func (o Opinion) String() string {
	switch o {
	case OpinionOther:
		return "O"
	case OpinionFarLeft:
		return "XL"
	case OpinionLeft:
		return "L"
	case OpinionCenter:
		return "C"
	case OpinionRight:
		return "R"
	case OpinionFarRight:
		return "XR"

	case OpinionBlank:
		return "B"
	case OpinionNull:
		return "N"
	case OpinionAbstention:
		return "A"

	default:
		return "_"
	}
}

func (o Opinion) Title() string {
	switch o {
	case OpinionOther:
		return "Autre"
	case OpinionFarLeft:
		return "Extrême gauche"
	case OpinionLeft:
		return "Gauche"
	case OpinionCenter:
		return "Centre"
	case OpinionRight:
		return "Droite"
	case OpinionFarRight:
		return "Extrême droite"

	case OpinionBlank:
		return "Blanc"
	case OpinionNull:
		return "Nul"
	case OpinionAbstention:
		return "Abstention"

	default:
		return "???"
	}
}
