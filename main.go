package main

import (
	_ "embed"
	"lfi/data-vote/common"
	"lfi/data-vote/render"
	"lfi/data-vote/votation"
	"lfi/data-vote/votation/legislative2024"
	"lfi/data-vote/votation/presidentielle2017"
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
		func(t *tool.Tool) []*common.Event { return tr(legislative2024.Fetch(t)) },
		func(t *tool.Tool) []*common.Event { return tr(ue2024.Parse(t)) },
		presidentielle2017.Fetch,
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
	case "", "Troyes", "Saint-Julien-les-Villas":
		return false
	default:
		return true
	}
}

func tr(stations []*votation.Station) (events []*common.Event) {
	events = make([]*common.Event, len(stations))
	for i, s := range stations {
		v := s.Votation[0]
		options := make([]common.Option, len(v.Result))
		for i, r := range v.Result {
			options[i] = common.Option{
				Result:   r.Result,
				Position: r.Position,
				Party:    r.Party,
				Opinion:  common.Opinion(r.Opinion),
				Name:     r.Name,
				Gender:   common.Gender(r.Gender),
			}
		}

		events[i] = &common.Event{
			Departement: common.Departement(s.Departement),
			City:        s.City,
			StationID:   s.CodeStation,

			VoteID:   v.Code,
			VoteName: v.Name,

			Register:   v.Register,
			Abstention: v.Abstention,
			Blank:      v.Blank,
			Null:       v.Null,

			Option: options,
		}
	}
	return
}
