package legislative2024

import (
	"encoding/csv"
	"fmt"
	"lfi/data-vote/votation"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strconv"
	"strings"
)

func Fetch(t *tool.Tool) []*votation.Station {
	body := t.Fetch(fetch.URL(url)).Body
	defer body.Close()

	r := csv.NewReader(body)
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		t.Error("csv.parse", "err", err.Error())
		return nil
	} else if strings.Join(lines[0], ";") != header {
		t.Error("wrong header")
		return nil
	}

	stations := make([]*votation.Station, 0, len(lines[1:]))
	for i, line := range lines[1:] {
		s, err := parseLine(line)
		if err != nil {
			t.Warn("parse line", "err", err.Error(), "line", i+2)
			continue
		}
		stations = append(stations, s)
	}

	return stations
}

func parseLine(line []string) (_ *votation.Station, err error) {
	r := votation.VotationResult{}
	r.Register = parseUint(&err, line[5])
	r.Abstention = parseUint(&err, line[8])
	r.Blank = parseUint(&err, line[13])
	r.Null = parseUint(&err, line[16])
	departement := votation.DepartementName2Const[line[1]]
	if err != nil {
		return nil, err
	} else if departement == 0 {
		return nil, fmt.Errorf("Unknow departement %q", line[1])
	}

	err = parseResults(&r, line[19:])
	if err != nil {
		return nil, err
	}

	return &votation.Station{
		Departement: departement,
		City:        line[3],
		CodeStation: line[4],
		Votation: []votation.Votation{
			{Name: voteName, Date: voteDate, Code: voteCode, VotationResult: r},
		},
	}, nil
}

func parseResults(r *votation.VotationResult, line []string) (err error) {
	if len(line) == 0 {
		return nil
	} else if len(line) < 9 {
		return fmt.Errorf("Expected at least 9 element, get only %d", len(line))
	} else if line[0] == "" {
		return nil
	}

	r.Result = append(r.Result, votation.Result{
		Option: &votation.Option{
			Position: parseUint(&err, line[0]),
			Party:    line[1],
			Opinion:  parseOpinion(&err, line[1]),
			Name:     line[3] + " " + line[2],
			Gender:   parseGender(&err, line[4]),
		},
		Result: parseUint(&err, line[5]),
	})
	if err != nil {
		return err
	}

	return parseResults(r, line[9:])
}

func parseUint(perr *error, s string) uint {
	if *perr != nil {
		return 0
	}
	u, err := strconv.ParseUint(s, 10, 30)
	if err != nil {
		*perr = err
		return 0
	}
	return uint(u)
}

func parseOpinion(perr *error, s string) votation.Opinion {
	if *perr != nil {
		return 0
	}
	// Source: https://www.vie-publique.fr/en-bref/285049-legislatives-2022-une-circulaire-dattribution-des-nuances-politiques
	// And a lot of empirism
	switch s {
	case "REG", "DIV":
		return votation.OpinionOther
	case "DXG", "EXG":
		return votation.OpinionFarLeft
	case "COM", "FI", "ECO", "UG", "VEC":
		return votation.OpinionLeft
	case "SOC", "RDG", "DVC", "DVG":
		return votation.OpinionCenter
	case "ENS", "LR", "UDI", "DVD", "DSV", "HOR":
		return votation.OpinionRight
	case "REC", "RN", "DXD", "EXD", "UXD":
		return votation.OpinionFarRight
	default:
		*perr = fmt.Errorf("Unknown opinion %q", s)
		return 0
	}
}

func parseGender(perr *error, s string) votation.Gender {
	if *perr != nil {
		return 0
	}
	switch s {
	case "MASCULIN":
		return votation.GenderMan
	case "FEMININ":
		return votation.GenderWoman
	default:
		*perr = fmt.Errorf("Sex: expected 'MASCULIN' or 'FEMININ', get %q", s)
		return 0
	}
}
