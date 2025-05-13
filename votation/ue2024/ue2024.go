package ue2024

import (
	"encoding/csv"
	"fmt"
	"lfi/data-vote/votation"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strconv"
	"strings"
)

func Parse(t *tool.Tool) []*votation.Station {
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

	bureaux := make([]*votation.Station, 0, len(lines[1:]))
	for i, line := range lines[1:] {
		bv, err := parseLine(line)
		if err != nil {
			t.Warn("parse line", "err", err.Error(), "line", i+2)
			continue
		}
		bureaux = append(bureaux, bv)
	}

	return bureaux
}

func parseLine(line []string) (_ *votation.Station, err error) {
	r := votation.VotationResult{}
	r.Register = parseUint(&err, line[7])
	r.Abstention = parseUint(&err, line[10])
	r.Blank = parseUint(&err, line[15])
	r.Null = parseUint(&err, line[18])
	departement := votation.DepartementName2Const[line[3]]
	if err != nil {
		return nil, err
	} else if departement == 0 {
		return nil, fmt.Errorf("Unknow departement %q", line[1])
	}

	r.Result = make([]votation.Result, 0, 38)
	if err := parseOption(&r, line[21:]); err != nil {
		return nil, err
	}

	return &votation.Station{
		Departement: departement,
		City:        line[5],
		CodeStation: strings.TrimLeft(line[6], "0"),
		Votation: []votation.Votation{
			{Name: voteName, Date: voteDate, VotationResult: r},
		},
	}, nil
}

func parseOption(r *votation.VotationResult, line []string) (err error) {
	if len(line) == 0 {
		return nil
	}

	result := parseUint(&err, line[4])
	if err != nil {
		return
	}

	r.Result = append(r.Result, votation.Result{
		Option: options[len(r.Result)],
		Result: result,
	})

	return parseOption(r, line[8:])
}

func parseUint(perr *error, s string) uint {
	if *perr != nil {
		return 0
	} else if s == "-1" {
		return 0
	}
	u, err := strconv.ParseUint(s, 10, 30)
	if err != nil {
		*perr = err
		return 0
	}
	return uint(u)
}
