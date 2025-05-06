package driver2024legislative

import (
	"encoding/csv"
	"fmt"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strconv"
	"strings"
)

const header = `Code département;Libellé département;Code commune;Libellé commune;Code BV;Inscrits;Votants;% Votants;Abstentions;% Abstentions;Exprimés;% Exprimés/inscrits;% Exprimés/votants;Blancs;% Blancs/inscrits;% Blancs/votants;Nuls;% Nuls/inscrits;% Nuls/votants;Numéro de panneau 1;Nuance candidat 1;Nom candidat 1;Prénom candidat 1;Sexe candidat 1;Voix 1;% Voix/inscrits 1;% Voix/exprimés 1;Elu 1;Numéro de panneau 2;Nuance candidat 2;Nom candidat 2;Prénom candidat 2;Sexe candidat 2;Voix 2;% Voix/inscrits 2;% Voix/exprimés 2;Elu 2;Numéro de panneau 3;Nuance candidat 3;Nom candidat 3;Prénom candidat 3;Sexe candidat 3;Voix 3;% Voix/inscrits 3;% Voix/exprimés 3;Elu 3;Numéro de panneau 4;Nuance candidat 4;Nom candidat 4;Prénom candidat 4;Sexe candidat 4;Voix 4;% Voix/inscrits 4;% Voix/exprimés 4;Elu 4;Numéro de panneau 5;Nuance candidat 5;Nom candidat 5;Prénom candidat 5;Sexe candidat 5;Voix 5;% Voix/inscrits 5;% Voix/exprimés 5;Elu 5;Numéro de panneau 6;Nuance candidat 6;Nom candidat 6;Prénom candidat 6;Sexe candidat 6;Voix 6;% Voix/inscrits 6;% Voix/exprimés 6;Elu 6;Numéro de panneau 7;Nuance candidat 7;Nom candidat 7;Prénom candidat 7;Sexe candidat 7;Voix 7;% Voix/inscrits 7;% Voix/exprimés 7;Elu 7;Numéro de panneau 8;Nuance candidat 8;Nom candidat 8;Prénom candidat 8;Sexe candidat 8;Voix 8;% Voix/inscrits 8;% Voix/exprimés 8;Elu 8;Numéro de panneau 9;Nuance candidat 9;Nom candidat 9;Prénom candidat 9;Sexe candidat 9;Voix 9;% Voix/inscrits 9;% Voix/exprimés 9;Elu 9;Numéro de panneau 10;Nuance candidat 10;Nom candidat 10;Prénom candidat 10;Sexe candidat 10;Voix 10;% Voix/inscrits 10;% Voix/exprimés 10;Elu 10;Numéro de panneau 11;Nuance candidat 11;Nom candidat 11;Prénom candidat 11;Sexe candidat 11;Voix 11;% Voix/inscrits 11;% Voix/exprimés 11;Elu 11;Numéro de panneau 12;Nuance candidat 12;Nom candidat 12;Prénom candidat 12;Sexe candidat 12;Voix 12;% Voix/inscrits 12;% Voix/exprimés 12;Elu 12;Numéro de panneau 13;Nuance candidat 13;Nom candidat 13;Prénom candidat 13;Sexe candidat 13;Voix 13;% Voix/inscrits 13;% Voix/exprimés 13;Elu 13;Numéro de panneau 14;Nuance candidat 14;Nom candidat 14;Prénom candidat 14;Sexe candidat 14;Voix 14;% Voix/inscrits 14;% Voix/exprimés 14;Elu 14;Numéro de panneau 15;Nuance candidat 15;Nom candidat 15;Prénom candidat 15;Sexe candidat 15;Voix 15;% Voix/inscrits 15;% Voix/exprimés 15;Elu 15;Numéro de panneau 16;Nuance candidat 16;Nom candidat 16;Prénom candidat 16;Sexe candidat 16;Voix 16;% Voix/inscrits 16;% Voix/exprimés 16;Elu 16;Numéro de panneau 17;Nuance candidat 17;Nom candidat 17;Prénom candidat 17;Sexe candidat 17;Voix 17;% Voix/inscrits 17;% Voix/exprimés 17;Elu 17;Numéro de panneau 18;Nuance candidat 18;Nom candidat 18;Prénom candidat 18;Sexe candidat 18;Voix 18;% Voix/inscrits 18;% Voix/exprimés 18;Elu 18;Numéro de panneau 19;Nuance candidat 19;Nom candidat 19;Prénom candidat 19;Sexe candidat 19;Voix 19;% Voix/inscrits 19;% Voix/exprimés 19;Elu 19`

const url = "https://static.data.gouv.fr/resources/elections-legislatives-des-30-juin-et-7-juillet-2024-resultats-definitifs-du-1er-tour/20240710-171445/resultats-definitifs-par-bureau-de-vote.csv"

func Parse(t *tool.Tool) {
	body := t.Fetch(fetch.URL(url)).Body
	defer body.Close()

	r := csv.NewReader(body)
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		t.Error("csv.parse", "err", err.Error())
		return
	} else if strings.Join(lines[0], ";") != header {
		t.Error("wrong header")
		return
	}

	bureaux := make([]*BureauVote, 0, len(lines[1:]))
	for i, line := range lines[1:] {
		bv, err := parseBureauVote(line)
		if err != nil {
			t.Warn("parse line", "err", err.Error(), "line", i+2)
		}
		bureaux = append(bureaux, bv)
	}

}

type BureauVote struct {
	DépartementCode string
	DépartementName string
	Commune         string
	CodeBV          string

	Inscrit     uint
	Abstentions uint
	Blancs      uint
	Nuls        uint

	Votes []Vote
}

type Vote struct {
	Panneau  int
	Nuance   string
	Nom      string
	Prénom   string
	EstFemme bool
	Voix     int
}

func parseBureauVote(line []string) (*BureauVote, error) {
	inscrit, err := strconv.ParseUint(line[5], 10, 30)
	if err != nil {
		return nil, fmt.Errorf("Parse inscrit %q: %w", line[5], err)
	}

	abstentions, err := strconv.ParseUint(line[8], 10, 30)
	if err != nil {
		return nil, fmt.Errorf("Parse abstentions %q: %w", line[8], err)
	}

	blancs, err := strconv.ParseUint(line[13], 10, 30)
	if err != nil {
		return nil, fmt.Errorf("Parse Blancs %q: %w", line[13], err)
	}

	nuls, err := strconv.ParseUint(line[16], 10, 30)
	if err != nil {
		return nil, fmt.Errorf("Parse inscrit %q: %w", line[16], err)
	}

	votes, err := parseVotes(line[19:], make([]Vote, 0, 19))
	if err != nil {
		return nil, err
	}

	return &BureauVote{
		DépartementCode: line[0],
		DépartementName: line[1],
		Commune:         line[3],
		CodeBV:          line[4],

		Inscrit:     uint(inscrit),
		Abstentions: uint(abstentions),
		Blancs:      uint(blancs),
		Nuls:        uint(nuls),

		Votes: votes,
	}, nil

	// 0: Code département
	// 1: Libellé département
	// 2: Code commune
	// 3: Libellé commune
	// 4: Code BV
	// 5: Inscrits

	// 6: Votants
	// 7: % Votants
	// 8: Abstentions
	// 9: % Abstentions

	// 10: Exprimés
	// 11: % Exprimés/inscrits
	// 12: % Exprimés/votants
	// 13: Blancs
	// 14: % Blancs/inscrits
	// 15: % Blancs/votants
	// 16: Nuls
	// 17: % Nuls/inscrits
	// 18: % Nuls/votants
}

// Parse the next line
func parseVotes(line []string, votes []Vote) ([]Vote, error) {
	if len(line) == 0 {
		return votes, nil
	} else if len(line) < 9 {
		return nil, fmt.Errorf("Expected at least 9 element, get only %d", len(line))
	} else if line[0] == "" {
		return votes, nil
	}

	panneau, err := strconv.ParseUint(line[0], 10, 5)
	if err != nil {
		return nil, fmt.Errorf("Get panneau num: %w", err)
	}

	estFemme := false
	switch line[4] {
	case "MASCULIN":
	case "FEMININ":
		estFemme = true
	default:
		return nil, fmt.Errorf("Sex: expected 'MASCULIN' or 'FEMININ', geet %q", line[4])
	}

	voix, err := strconv.ParseUint(line[5], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("Get voix number: %w", err)
	}

	votes = append(votes, Vote{
		Panneau:  int(panneau),
		Nuance:   line[1],
		Nom:      line[2],
		Prénom:   line[3],
		EstFemme: estFemme,
		Voix:     int(voix),
	})

	return parseVotes(line[9:], votes)

	// 0: Numéro de panneau N
	// 1: Nuance candidat N
	// 2: Nom candidat N
	// 3: Prénom candidat N
	// 4: Sexe candidat N
	// 5: Voix N
	// 6: % Voix/inscrits N
	// 7: % Voix/exprimés N
	// 8: Elu N
}
