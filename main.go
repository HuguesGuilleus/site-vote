package main

import (
	_ "embed"
	"lfi/data-vote/render"
	"lfi/data-vote/votation"
	"lfi/data-vote/votation/legislative2024"
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

	stations := legislative2024.Fetch(t)
	for _, s := range stations {
		if s.Departement != votation.DepartementAube {
			continue
		}
		render.RenderStation(t, s)
	}

	// r := t.Fetch(fetch.URL("https://static.data.gouv.fr/resources/resultats-des-elections-europeennes-du-9-juin-2024/20240613-154804/resultats-definitifs-par-bureau-de-vote.csv"))
	// io.Copy(os.Stdout, io.LimitReader(r.Body, 4096*5))
}
