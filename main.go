package main

import (
	_ "embed"

	"github.com/HuguesGuilleus/site-vote/common"
	"github.com/HuguesGuilleus/site-vote/render"
	"github.com/HuguesGuilleus/site-vote/votation/legislative2012"
	"github.com/HuguesGuilleus/site-vote/votation/legislative2017"
	"github.com/HuguesGuilleus/site-vote/votation/legislative2022"
	"github.com/HuguesGuilleus/site-vote/votation/legislative2024"
	"github.com/HuguesGuilleus/site-vote/votation/municipale2014"
	"github.com/HuguesGuilleus/site-vote/votation/municipale2020"
	"github.com/HuguesGuilleus/site-vote/votation/presidentielle2012"
	"github.com/HuguesGuilleus/site-vote/votation/presidentielle2017"
	"github.com/HuguesGuilleus/site-vote/votation/presidentielle2022"
	"github.com/HuguesGuilleus/site-vote/votation/ue2024"
	"github.com/HuguesGuilleus/sniffle/tool"
)

//go:embed style.css
var style []byte

//go:embed favicon.ico
var favicon []byte

func main() {
	t := tool.New(tool.CLI(nil))

	t.Info("fetch ...")
	events := common.Call(t,
		municipale2020.Fetch,
		municipale2014.Fetch,

		legislative2024.Fetch,
		legislative2022.Fetch,
		legislative2017.Fetch,
		legislative2012.Fetch,

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
	for z := range common.ByLegislativeDisctrict(events, skip) {
		render.RenderLegislativeDisctrict(t, z)
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
