package legislative2017

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"strings"

	"github.com/HuguesGuilleus/sniffle/tool"
)

const (
	url = "https://static.data.gouv.fr/resources/elections-legislatives-des-11-et-18-juin-2017-resultats-par-bureaux-de-vote-du-1er-tour/20170613-100441/Leg_2017_Resultats_BVT_T1_c.txt"

	header = "Code du département;Libellé du département;Code de la circonscription;Libellé de la circonscription;Code de la commune;Libellé de la commune;Code du b.vote;Inscrits;Abstentions;% Abs/Ins;Votants;% Vot/Ins;Blancs;% Blancs/Ins;% Blancs/Vot;Nuls;% Nuls/Ins;% Nuls/Vot;Exprimés;% Exp/Ins;% Exp/Vot;N°Panneau;Sexe;Nom;Prénom;Nuance;Voix;% Voix/Ins;% Voix/Exp"

	voteID   = "2017_06_11_L"
	voteName = "2017-06-11 Législatives"
)

func Fetch(t *tool.Tool) (events []*common.Event) {
	lines := csvtool.FetchWeirdCSV(t, url, header)
	events = make([]*common.Event, len(lines))
	for i, line := range lines {
		events[i] = &common.Event{
			Departement: common.DepartementCode2Const[line[0]],
			City:        line[5],
			StationID:   strings.TrimLeft(line[6], "0"),
			District:    strings.TrimLeft(line[2], "0"),

			VoteID:   voteID,
			VoteName: voteName,

			Register:   csvtool.ParseUint(line[7]),
			Abstention: csvtool.ParseUint(line[8]),
			Blank:      csvtool.ParseUint(line[12]),
			Null:       csvtool.ParseUint(line[15]),

			Option: parseOption(line[21:], make([]common.Option, 0)),
		}
	}
	return
}

func parseOption(line []string, options []common.Option) []common.Option {
	if len(line) == 0 {
		return options
	}

	gender := common.GenderMan
	if line[1] == "F" {
		gender = common.GenderWoman
	}

	options = append(options, common.Option{
		Result:   csvtool.ParseUint(line[5]),
		Position: csvtool.ParseUint(line[0]),
		Party:    line[4],
		Opinion:  parseOpinion(line[4]),
		Name:     line[3] + " " + line[2],
		Gender:   gender,
	})

	return parseOption(line[8:], options)
}

func parseOpinion(party string) common.Opinion {
	switch party {
	case "EXG":
		return common.OpinionFarLeft
	case "FI", "DVG", "COM":
		return common.OpinionLeft
	case "SOC", "ECO", "REM", "RDG":
		return common.OpinionCenter
	case "UDI", "LR", "DVD", "MDM":
		return common.OpinionRight
	case "FN", "EXD", "DLF":
		return common.OpinionFarRight
	case "DIV", "REG":
		return common.OpinionOther
	default:
		panic("Unknown party: " + party)
	}
}
