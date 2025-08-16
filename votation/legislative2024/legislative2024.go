package legislative2024

import (
	"strings"

	"github.com/HuguesGuilleus/site-vote/common"
	"github.com/HuguesGuilleus/site-vote/common/csvtool"
	"github.com/HuguesGuilleus/sniffle/tool"
)

type districtByNames []map[string]string

func Fetch(t *tool.Tool) []*common.Event {
	lines := csvtool.FetchCSV(t, url, header)
	events := make([]*common.Event, len(lines))
	for i, line := range lines {
		events[i] = parseEvent(line)
	}

	districts := findDistrict(t)
	for _, event := range events {
		districtKey := strings.ToLower("1 " + event.Option[0].Name)
		event.District = districts[event.Departement][districtKey]
	}

	return events
}

func parseEvent(line []string) *common.Event {
	return &common.Event{
		Departement: common.DepartementName2Const[line[1]],
		City:        line[3],
		StationID:   line[4],

		VoteID:   voteID,
		VoteName: voteName,

		Register:   csvtool.ParseUint(line[5]),
		Abstention: csvtool.ParseUint(line[8]),
		Blank:      csvtool.ParseUint(line[13]),
		Null:       csvtool.ParseUint(line[16]),

		Option: parseOption(line[19:], make([]common.Option, 0, 19)),
	}
}

func findDistrict(t *tool.Tool) (districts districtByNames) {
	districts = make(districtByNames, common.DepartementLen)
	for d := range districts {
		districts[d] = make(map[string]string)
	}
	for _, candidate := range csvtool.FetchCSV(t, urlCirco, headerCirco) {
		d := common.DepartementCode2Const[candidate[0]]
		districtKey := strings.ToLower("" +
			candidate[4] + " " +
			candidate[7] + " " +
			candidate[8],
		)
		districts[d][districtKey] = strings.TrimLeft(strings.TrimPrefix(candidate[2], candidate[0]), "0")
	}
	return
}

func parseOption(line []string, options []common.Option) []common.Option {
	if len(line) == 0 || line[0] == "" {
		return options
	}

	options = append(options, common.Option{
		Result:   csvtool.ParseUint(line[5]),
		Position: uint(len(options)) + 1,
		Party:    line[1],
		Opinion:  parseOpinion(line[1]),
		Name:     line[2] + " " + line[3],
		Gender:   parseGender(line[4]),
	})

	return parseOption(line[9:], options)
}

func parseOpinion(s string) common.Opinion {
	// Source: https://www.vie-publique.fr/en-bref/285049-legislatives-2022-une-circulaire-dattribution-des-nuances-politiques
	// And a lot of empirism
	switch s {
	case "REG", "DIV":
		return common.OpinionOther
	case "DXG", "EXG":
		return common.OpinionFarLeft
	case "COM", "FI", "ECO", "UG", "VEC":
		return common.OpinionLeft
	case "SOC", "RDG", "DVC", "DVG":
		return common.OpinionCenter
	case "ENS", "LR", "UDI", "DVD", "DSV", "HOR":
		return common.OpinionRight
	case "REC", "RN", "DXD", "EXD", "UXD":
		return common.OpinionFarRight
	default:
		panic("Unknwo opinion: " + s)
	}
}

func parseGender(s string) common.Gender {
	switch s {
	case "MASCULIN":
		return common.GenderMan
	case "FEMININ":
		return common.GenderWoman
	default:
		panic("Unknown gender: " + s)
	}
}
