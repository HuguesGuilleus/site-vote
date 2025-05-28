package presidentielle2022

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"sniffle/tool"
	"strings"
)

const (
	url = "https://static.data.gouv.fr/resources/election-presidentielle-des-10-et-24-avril-2022-resultats-definitifs-du-1er-tour/20220414-152542/resultats-par-niveau-burvot-t1-france-entiere.txt"

	header = `Code du département;Libellé du département;Code de la circonscription;Libellé de la circonscription;Code de la commune;Libellé de la commune;Code du b.vote;Inscrits;Abstentions;% Abs/Ins;Votants;% Vot/Ins;Blancs;% Blancs/Ins;% Blancs/Vot;Nuls;% Nuls/Ins;% Nuls/Vot;Exprimés;% Exp/Ins;% Exp/Vot;N°Panneau;Sexe;Nom;Prénom;Voix;% Voix/Ins;% Voix/Exp`

	voteID   = "2022_03_03_P"
	voteName = "2022-03-03 Présidentielle"
)

func Fetch(t *tool.Tool) []*common.Event {
	lines := csvtool.FetchWeirdCSV(t, url, header)
	events := make([]*common.Event, len(lines))
	for i, line := range lines {
		events[i] = parseEvent(line)
	}
	return events
}

func parseEvent(line []string) *common.Event {
	options := constOptions.Clone()
	for i := range options {
		options[i].Result = csvtool.ParseUint(line[21+i*7+4])
	}

	return &common.Event{
		Departement: common.DepartementName2Const[line[1]],
		City:        line[5],
		StationID:   strings.TrimLeft(line[6], "0"),

		VoteID:   voteID,
		VoteName: voteName,

		Register:   csvtool.ParseUint(line[7]),
		Abstention: csvtool.ParseUint(line[8]),
		Blank:      csvtool.ParseUint(line[12]),
		Null:       csvtool.ParseUint(line[15]),

		Option: options,
	}
}

var constOptions = common.ConstOptions(
	`XL	LO	F	Nathalie ARTHAUD`,
	`L	PCF	M	Fabien ROUSSEL`,
	`R	EM	M	Emmanuel	MACRON`,
	`R	RES	M	Jean LASSALLE`,
	`XR	RN	F	Marine LE PEN`,
	`XR	REC	M	Éric ZEMMOUR`,
	`L	LFI	M	Jean-Luc MÉLENCHON`,
	`C	PS	F	Anne HIDALGO`,
	`L	EELV	M	Yannick JADOT`,
	`R	LR	F	Valérie PÉCRESSE`,
	`XL	NPA	M	Philippe	POUTOU`,
	`XR	DLF	M	Nicolas DUPONT-AIGNAN`,
)
