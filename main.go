package main

import (
	_ "embed"
	"lfi/data-vote/common"
	"lfi/data-vote/render"
	"lfi/data-vote/votation/legislative2024"
	"lfi/data-vote/votation/presidentielle2012"
	"lfi/data-vote/votation/presidentielle2017"
	"lfi/data-vote/votation/presidentielle2022"
	"lfi/data-vote/votation/ue2024"
	"sniffle/tool"
)

//go:embed style.css
var style []byte

//go:embed favicon.ico
var favicon []byte

func main() {
	t := tool.New(tool.CLI(nil))

	t.Info("fetch ...")
	events := common.Call(t,
		legislative2024.Fetch,

		presidentielle2022.Fetch,
		presidentielle2017.Fetch,
		presidentielle2012.Fetch,

		ue2024.Fetch,
	)

	t.Info("render ...")
	t.WriteFile("/style.css", style)
	t.WriteFile("/favicon.ico", favicon)
	render.RenderFrance(t, common.AllFrance(events))
	for z := range common.ByDepartement(events, skip) {
		render.RenderDepartement(t, z)
	}
	for z := range common.ByCity(events, skip) {
		render.RenderCity(t, z)
	}
	for z := range common.ByStation(events, skip) {
		render.RenderStation(t, z)
	}
}

func skip(d common.Departement, city string) bool {
	switch d {
	case common.DepartementAube:
	default:
		return true
	}
	switch city {
	case "", "La Chapelle-Saint-Luc", "Saint-Julien-les-Villas", "Sainte-Savine", "Troyes":
		return false
	default:
		return true
	}
}
