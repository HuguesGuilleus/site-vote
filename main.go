package main

import (
	_ "embed"
	"lfi/data-vote/render"
	"lfi/data-vote/votation"
	"lfi/data-vote/votation/legislative2024"
	"lfi/data-vote/votation/ue2024"
	"sniffle/tool"
)

//go:embed style.css
var style []byte

func main() {
	t := tool.New(tool.CLI(nil))

	t.WriteFile("/vote/style.css", style)

	// bureaux := driver2024legislative.Parse(t)
	// for _, bv := range bureaux {
	// 	if bv.DépartementCode != "10" {
	// 		continue
	// 	}
	// 	render.RenderBV(t, bv)
	// }

	stations := votation.MergeStation(
		legislative2024.Fetch(t),
		ue2024.Parse(t),
	)

	for _, s := range stations {
		if s.Departement != votation.DepartementAube {
			continue
		}
		render.RenderStation(t, s)
	}
}
