package municipale2014

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"maps"
	"slices"
	"strings"

	"github.com/HuguesGuilleus/sniffle/tool"
)

const (
	url = "https://www.data.gouv.fr/s/resources/elections-municipales-2014-resultats-par-bureaux-de-vote/20150925-122128/MN14_Bvot_T1T2.txt"

	voteID   = "2014_03_23_M"
	voteName = "2014-03-23 Municipales"
)

func Fetch(t *tool.Tool) []*common.Event {
	type Location struct {
		Departement common.Departement
		City        string
		StationID   string
	}
	locations := make(map[Location]*common.Event)

	lines := csvtool.FetchWeirdCSV(t, url, "")
	for _, line := range lines {
		if strings.HasPrefix(line[0], "--") {
			continue
		} else if line[0] == "2" {
			// No first tour after second
			break
		}

		l := Location{
			Departement: common.DepartementCode2Const[line[1]],
			City:        line[3],
			StationID:   strings.TrimLeft(line[4], "0"),
		}
		e := locations[l]
		if e == nil {
			register := csvtool.ParseUint(line[5])
			voter := csvtool.ParseUint(line[6])
			valid := csvtool.ParseUint(line[7])
			e = &common.Event{
				Departement: l.Departement,
				City:        l.City,
				StationID:   l.StationID,

				VoteID:   voteID,
				VoteName: voteName,

				Register:   register,
				Abstention: register - voter,
				Null:       voter - valid,

				Option: make([]common.Option, 0),
			}
			locations[l] = e
		}

		e.Option = append(e.Option, common.Option{
			Result:   csvtool.ParseUint(line[12]),
			Position: csvtool.ParseUint(line[8]),
			Party:    line[11],
			Opinion:  parseOpinion(line[11]),
			Name:     line[10] + " " + line[9],
		})
	}

	events := slices.AppendSeq(
		make([]*common.Event, 0, len(locations)),
		maps.Values(locations),
	)
	common.SetSplitVoting(events)
	return events
}

func parseOpinion(party string) common.Opinion {
	// Source: https://www.archives-resultats-elections.interieur.gouv.fr/resultats/MN2014/nuances.php
	switch party {
	case "LEXG":
		return common.OpinionFarLeft
	case "LCOM", "LDVG", "LFG", "LPG", "LSOC", "LUG", "LVEC":
		return common.OpinionLeft
	case "LMDM", "LUC":
		return common.OpinionCenter
	case "LDVD", "LUD", "LUDI", "LUMP":
		return common.OpinionRight
	case "LEXD", "LFN":
		return common.OpinionFarRight
	case "LDIV", "NC":
		return common.OpinionOther
	default:
		panic("unknown party: " + party)
	}
}
