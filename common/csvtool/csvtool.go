package csvtool

import (
	"encoding/csv"
	"io"
	"sniffle/tool"
	"sniffle/tool/fetch"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func FetchCSV(t *tool.Tool, url, header string) [][]string {
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

func FetchWeirdCSV(t *tool.Tool, url, header string) (lines [][]string) {
	r := t.Fetch(fetch.URL(url))
	defer r.Body.Close()
	data, err := io.ReadAll(charmap.ISO8859_1.NewDecoder().Reader(r.Body))
	if err != nil {
		t.Error("read all body fail", "err", err.Error())
		return nil
	}

	lineData := strings.Split(string(data), "\r\n")
	if header != "" {
		if lineData[0] != header {
			t.Error("wrong header")
			return nil
		}
		lineData = lineData[1:]
	}

	lines = make([][]string, len(lineData))
	for i, line := range lineData {
		lines[i] = strings.Split(line, ";")
	}

	if i := len(lines) - 1; len(lines[i]) == 1 && lines[i][0] == "" {
		lines = lines[:i]
	}

	return
}
