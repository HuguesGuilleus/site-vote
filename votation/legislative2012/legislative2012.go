package legislative2012

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"maps"
	"slices"
	"strings"

	"github.com/HuguesGuilleus/sniffle/tool"
)

const (
	url = "https://static.data.gouv.fr/resources/elections-legislatives-2012-resultats-par-bureaux-de-vote/20150925-103435/LG12_Bvot_T1T2.txt"

	voteID   = "2012_06_10_L"
	voteName = "2012-06-10 Législatives"
)

func Fetch(t *tool.Tool) []*common.Event {
	type Location struct {
		Departement common.Departement
		City        string
		StationID   string
	}
	locations := make(map[Location]*common.Event)

	for _, line := range csvtool.FetchWeirdCSV(t, url, "") {
		if strings.HasPrefix(line[0], "--") {
			continue
		} else if line[0] == "2" {
			break
		}

		location := Location{
			Departement: common.DepartementCode2Const[line[1]],
			City:        line[3],
			StationID:   strings.TrimLeft(line[6], "0"),
		}
		event, ok := locations[location]
		if !ok {
			register := csvtool.ParseUint(line[7])
			voters := csvtool.ParseUint(line[8])
			valids := csvtool.ParseUint(line[9])
			event = &common.Event{
				Departement: common.DepartementCode2Const[line[1]],
				City:        line[3],
				StationID:   strings.TrimLeft(line[6], "0"),
				District:    strings.TrimLeft(line[4], "0"),

				VoteID:   voteID,
				VoteName: voteName,

				Register:   register,
				Abstention: register - voters,
				Null:       voters - valids,

				Option: make([]common.Option, 0),
			}
			locations[location] = event
		}

		event.Option = append(event.Option, common.Option{
			Result:   csvtool.ParseUint(line[14]),
			Position: csvtool.ParseUint(line[10]),
			Party:    line[13],
			Opinion:  parseParty(line[13]),
			Name:     line[12] + " " + line[11],
		})
	}

	return slices.AppendSeq(
		make([]*common.Event, 0, len(locations)),
		maps.Values(locations),
	)
}

func parseParty(party string) common.Opinion {
	// Source: https://www.archives-resultats-elections.interieur.gouv.fr/resultats/LG2012/FE.php
	switch party {
	case "EXG":
		return common.OpinionFarLeft
	case "DVG", "ECO", "FG", "RDG", "SOC", "VEC":
		return common.OpinionLeft
	case "CEN", "ALLI", "PRV":
		return common.OpinionCenter
	case "DVD", "UMP", "NCE":
		return common.OpinionRight
	case "EXD", "FN":
		return common.OpinionFarRight
	case "AUT", "REG":
		return common.OpinionOther
	default:
		panic("unknown party: " + party)
	}
}
