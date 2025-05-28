package presidentielle2017

import (
	"io"
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const url = "https://static.data.gouv.fr/resources/election-presidentielle-des-23-avril-et-7-mai-2017-resultats-definitifs-du-1er-tour-par-bureaux-de-vote/20170427-100955/PR17_BVot_T1_FE.txt"

const header = `Code du département;Libellé du département;Code de la circonscription;Libellé de la circonscription;Code de la commune;Libellé de la commune;Code du b.vote;Inscrits;Abstentions;% Abs/Ins;Votants;% Vot/Ins;Blancs;% Blancs/Ins;% Blancs/Vot;Nuls;% Nuls/Ins;% Nuls/Vot;Exprimés;% Exp/Ins;% Exp/Vot;N°Panneau;Sexe;Nom;Prénom;Voix;% Voix/Ins;% Voix/Exp`

const (
	voteID   = "2017_04_23_P"
	voteName = "2017-04-23 Présidentielle"
)

func Fetch(t *tool.Tool) (events []*common.Event) {
	lines := globalParse(t)
	events = make([]*common.Event, 0, len(lines))
	for _, line := range lines {
		events = append(events, parseEvent(line))
	}
	return
}

func globalParse(t *tool.Tool) (lines [][]string) {
	r := t.Fetch(fetch.URL(url))
	defer r.Body.Close()
	data, err := io.ReadAll(charmap.ISO8859_1.NewDecoder().Reader(r.Body))
	if err != nil {
		t.Error("read all body fail", "err", err.Error())
		return nil
	}

	lineData := strings.Split(string(data), "\r\n")
	if lineData[0] != header {
		t.Error("wrong header")
		return nil
	}

	for _, line := range lineData[1:] {
		lines = append(lines, strings.Split(line, ";"))
	}
	lines = lines[:len(lines)-1]

	return
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
	"XR	DLF	M	Nicolas DUPONT-AIGNAN",
	"XR	RN	F	Marine LE PEN",
	"R	EM	M	Emmanuel MACRON",
	"L	PS	M	Benoît HAMON",
	"XL	LO	F	Nathalie ARTHAUD",
	"XL	NPA	M	Philippe POUTOU",
	"XR	SP	M	Jacques CHEMINADE",
	"R	RE	M	Jean LASSALLE",
	"L	LFI	M	Jean-Luc MÉLENCHON",
	"XR	UPR	M	François ASSELINEAU",
	"R	LR	M	François FILLON",
)
