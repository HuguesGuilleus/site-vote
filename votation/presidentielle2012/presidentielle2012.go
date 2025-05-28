package presidentielle2012

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"slices"
	"sniffle/tool"
	"strings"
)

const (
	url = "https://static.data.gouv.fr/resources/election-presidentielle-2012-resultats-par-bureaux-de-vote-1/20150925-102751/PR12_Bvot_T1T2.txt"

	voteID   = "2012_04_22_P"
	voteName = "2012-04-22 Présidentielle"
)

func Fetch(t *tool.Tool) (events []*common.Event) {
	lines := slices.DeleteFunc(
		csvtool.FetchWeirdCSV(t, url, ""),
		func(line []string) bool { return line[0] == "2" },
	)
	events = make([]*common.Event, 0, len(lines)/len(constOptions))
	for lines := range slices.Chunk(lines, len(constOptions)) {
		register := csvtool.ParseUint(lines[0][7])
		nbVote := csvtool.ParseUint(lines[0][8])
		nbValid := csvtool.ParseUint(lines[0][9])

		options := constOptions.Clone()
		for _, line := range lines {
			options[csvtool.ParseUint(line[10])-2].Result = csvtool.ParseUint(line[14])
		}

		events = append(events, &common.Event{
			Departement: common.DepartementCode2Const[lines[0][1]],
			City:        lines[0][3],
			StationID:   strings.TrimLeft(lines[0][6], "0"),

			VoteID:   voteID,
			VoteName: voteName,

			Register:   register,
			Abstention: register - nbVote,
			Null:       nbVote - nbValid,

			Option: options,
		})
	}
	return events
}

var constOptions = common.ConstOptions(
	"",
	"L	EELV	F	Eva JOLY",
	"XR	FN	F	Marine LE PEN",
	"R	UMP	M	Nicolas SARKOZY",
	"L	FG	M	Jean-Luc MÉLENCHON",
	"XL	NPA	M	Philippe POUTOU",
	"XL	LO	F	Nathalie ARTHAUD",
	"XR	SP	M	Jacques CHEMINADE",
	"R	MoDem	M	François BAYROU",
	"XR	DLR	M	Nicolas DUPONT-AIGNAN",
	"C	PS	M	François HOLLANDE",
)
