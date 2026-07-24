package municipale2026

import (
	"fmt"
	"strings"

	"github.com/HuguesGuilleus/site-vote/common"
	"github.com/HuguesGuilleus/site-vote/common/csvtool"
	"github.com/HuguesGuilleus/sniffle/tool"
)

const (
	url = "https://static.data.gouv.fr/resources/elections-municipales-2026-resultats-du-premier-tour/20260320-164249/municipales-2026-resultats-bv-par-communes-2026-03-20.csv"

	voteID   = "2026_03_15_M"
	voteName = "2020-03-15 Municipales"
)

func Fetch(t *tool.Tool) (events []*common.Event) {
	lines := csvtool.FetchCSV(t, url, "")[1:]
	events = make([]*common.Event, 0, len(lines))

	for _, line := range lines {
		// Because negative values in general columns
		switch line[3] {
		case `Saint-Cyr-du-Gault`,
			`Noroy`,
			`Sempigny`,
			`Le Poët-en-Percip`,
			`Léglantiers`,
			`Le Mesnil-sur-Bulles`:
			continue
		}

		events = append(events, &common.Event{
			Departement: common.DepartementCode2Const[line[0]],
			City:        line[3],
			StationID:   strings.TrimLeft(line[4], "0"),

			VoteID:   voteID,
			VoteName: voteName,

			Register:   csvtool.ParseUint(line[5]),
			Abstention: csvtool.ParseUint(line[8]),
			Blank:      csvtool.ParseUint(line[13]),
			Null:       csvtool.ParseUint(line[16]),

			Option: parseOption(line[19:], make([]common.Option, 0)),
		})
	}

	common.SetSplitVoting(events)
	return
}

func parseOption(rest []string, options []common.Option) []common.Option {
	if len(rest) == 0 || rest[0] == "" {
		return options
	}

	opinion := common.Opinion(0)
	switch rest[4] {
	case "LEXG":
		opinion = common.OpinionFarLeft
	case "LDVG", "LUG", "LECO", "LFI", "LCOM", "LVEC":
		opinion = common.OpinionLeft
	case "LSOC", "LDVC", "LUC", "LMDM":
		opinion = common.OpinionCenter
	case "LDVD", "LR", "LLR", "LUDI", "LUDR", "LUD", "LREN", "LHOR":
		opinion = common.OpinionRight
	case "LRN", "LEXD", "LDSV", "LREC", "LUXD":
		opinion = common.OpinionFarRight
	case "LREG", "LDIV", "":
		opinion = common.OpinionOther
	default:
		fmt.Printf("options: %#+v\n", rest[4])
	}

	options = append(options, common.Option{
		Result:   csvtool.ParseUint(rest[7]),
		Position: csvtool.ParseUint(rest[0]),
		Party:    rest[4],
		Opinion:  opinion,
		Name:     rest[5],

		Gender: common.GenderMan,
	})

	return parseOption(rest[13:], options)
}
