package ue2024

import (
	"lfi/data-vote/common"
	"lfi/data-vote/common/csvtool"
	"sniffle/tool"
	"strings"
)

func Fetch(t *tool.Tool) []*common.Event {
	lines := csvtool.FetchCSV(t, url, header)
	events := make([]*common.Event, len(lines))
	for i, line := range lines {
		events[i] = parseEvent(line)
	}
	return events
}

func parseEvent(line []string) *common.Event {
	options := constOptions.Clone()
	for i := range options {
		options[i].Result = csvtool.ParseUint(line[21+i*8+4])
	}

	return &common.Event{
		Departement: common.DepartementName2Const[line[3]],
		City:        line[5],
		StationID:   strings.TrimLeft(line[6], "0"),

		VoteID:   voteID,
		VoteName: voteName,

		Register:   csvtool.ParseUint(line[7]),
		Abstention: csvtool.ParseUint(line[10]),
		Blank:      csvtool.ParseUint(line[15]),
		Null:       csvtool.ParseUint(line[18]),

		Option: options,
	}
}
