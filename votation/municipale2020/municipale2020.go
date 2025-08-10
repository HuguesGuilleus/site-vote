package municipale2020

import (
	"cmp"
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"slices"

	"github.com/HuguesGuilleus/sniffle/tool"
)

const (
	url = "https://static.data.gouv.fr/resources/elections-municipales-2020-resultats/20200525-133745/2020-05-18-resultats-par-niveau-burvot-t1-france-entiere.txt"

	voteID   = "2020_03_15_M"
	voteName = "2020-03-15 Municipales"
)

func Fetch(t *tool.Tool) (events []*common.Event) {
	lines := csvtool.FetchWeirdTab(t, url, "")[1:]
	events = make([]*common.Event, len(lines))
	for i, line := range lines {
		if line[0] == "24" && line[3] == "Le Lardin-Saint-Lazare" {
			line[24] += line[25]
			line = slices.Delete(line, 25, 26)
		}
		if line[0] == "29" && line[3] == "Brasparts" {
			line[24] += line[25]
			line = slices.Delete(line, 25, 26)
		}

		events[i] = &common.Event{
			Departement: common.DepartementCode2Const[line[0]],
			City:        line[3],
			StationID:   line[4],

			VoteID:   voteID,
			VoteName: voteName,

			Register:   csvtool.ParseUint(line[5]),
			Abstention: csvtool.ParseUint(line[6]),
			Blank:      csvtool.ParseUint(line[10]),
			Null:       csvtool.ParseUint(line[13]),

			Option: parseOption(line[19:], make([]common.Option, 0)),
		}
	}
	common.SetSplitVoting(events)
	return
}

func parseOption(line []string, options []common.Option) []common.Option {
	if len(line) == 0 || line[0] == "" {
		return options
	}

	gender := common.GenderMan
	if line[2] == "F" {
		gender = common.GenderWoman
	}

	options = append(options, common.Option{
		Result:   csvtool.ParseUint(line[6]),
		Position: csvtool.ParseUint(line[0]),
		Party:    line[1],
		Opinion:  parseOpinion(line[1]),
		Name:     "(" + cmp.Or(line[5], "~") + ") " + line[4] + " " + line[3],
		Gender:   gender,
	})

	return parseOption(line[9:], options)
}

func parseOpinion(party string) common.Opinion {
	// Source: https://www.archives-resultats-elections.interieur.gouv.fr/resultats/municipales-2020/nuances.php
	switch party {
	case "LEXG":
		return common.OpinionFarLeft
	case "LCOM", "LFI", "LSOC", "LRDG", "LDVG", "LUG", "LVEC", "LECO":
		return common.OpinionLeft
	case "LUC", "LDVC":
		return common.OpinionCenter
	case "LREM", "LMDM", "LUDI", "LLR", "LUD", "LDVD":
		return common.OpinionRight
	case "LDLF", "LRN", "LEXD":
		return common.OpinionFarRight
	case "LDIV", "LREG", "LGJ", "LNC", "NC":
		return common.OpinionOther
	default:
		panic("unknown party: " + party)
	}
}
