package csvtool

import (
	"encoding/csv"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strings"
)

func Fetch(t *tool.Tool, url, header string) [][]string {
	response := t.Fetch(fetch.URL(url))
	defer response.Body.Close()
	if response.Status != 200 {
		t.Error("Wrong HTTP status",
			"status", response.Status,
			"url", url,
		)
		return nil
	}

	r := csv.NewReader(response.Body)
	r.Comma = ';'
	lines, err := r.ReadAll()
	if err != nil {
		t.Error("csv.parse", "err", err.Error())
		return nil
	} else if strings.Join(lines[0], ";") != header {
		t.Error("wrong header")
		return nil
	}

	return lines[1:]
}
