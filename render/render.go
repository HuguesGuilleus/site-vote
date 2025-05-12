package render

import (
	"cmp"
	"fmt"
	"lfi/data-vote/driver2024legislative"
	. "lfi/data-vote/driver2024legislative"
	"slices"
	"sniffle/tool"
	"sniffle/tool/render"
)

func RenderBV(t *tool.Tool, bv *BureauVote) {
	m := merge{}
	m.mergeBureauVote(bv)

	slices.SortStableFunc(bv.Votes, func(a, b Vote) int { return cmp.Compare(a.Opinion, b.Opinion) })

	t.WriteFile(fmt.Sprintf("/vote/%s/%s/%s.html", bv.DépartementCode, bv.Commune, bv.CodeBV), render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=/vote/style.css>`),
			render.N("title", bv.Commune, ": ", bv.CodeBV),
		),
		render.N("body",
			render.N("nav", "..."),
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../../..").N("vote"),
					render.Na("a.headerBlock", "href", "..").N(bv.DépartementName),
					render.Na("a.headerBlock", "href", ".").N(bv.Commune),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Bureau:", bv.CodeBV),
				),
			),
			render.N("main",
				render.N("h1", "Législatives anticipées 2024"),
				render.N("div",
					render.N("div.bar", render.S2(m.Votes[:], "", func(o int, voices uint) render.Node {
						if voices == 0 || o == int(OpinionAbstention) {
							return render.Z
						}
						return render.Na("div.option",
							"data-opinion", Opinion(o).String()).
							A("style", fmt.Sprintf("width:%d%%", percent(voices, m.Register))).N()
					})),

					render.N("table",
						render.N("tr",
							render.N("th", "Voix"),
							render.N("th", "%"),
							render.N("th", "Nuance"),
							render.N("th", "Liste"),
						),
						render.S(bv.Votes, "", func(v Vote) render.Node {
							return render.N("tr",
								render.N("td.r", v.Voix),
								render.N("td.r", percent(v.Voix, m.Register), "%"),
								render.N("td",
									render.Na("div.copinion", "data-opinion", v.Opinion.String()),
									v.Nuance),
								render.N("td",
									"[", v.Panneau, "] ",
									v.Nom, " ", v.Prénom,
									render.IfSAny(v.EstFemme, " (elle)", " (il)"),
								),
							)
						}),
						render.N("tr",
							render.N("td.r", m.Votes[OpinionBlank]),
							render.N("td.r", percent(m.Votes[OpinionBlank], m.Register), "%"),
							render.N("td",
								render.Na("div.copinion", "data-opinion", OpinionBlank.String()),
								"Blanc"),
						),
						render.N("tr",
							render.N("td.r", m.Votes[OpinionNull]),
							render.N("td.r", percent(m.Votes[OpinionNull], m.Register), "%"),
							render.N("td",
								render.Na("div.copinion", "data-opinion", OpinionNull.String()),
								"Nul"),
						),
						render.N("tr",
							render.N("td.r", m.Votes[OpinionAbstention]),
							render.N("td.r", percent(m.Votes[OpinionAbstention], m.Register), "%"),
							render.N("td",
								render.Na("div.copinion", "data-opinion", OpinionAbstention.String()),
								"Abstention"),
						),
						render.N("tr",
							render.N("td.r", m.Register),
							render.N("td"),
							render.N("td", "Total"),
						),
					),
				),
			),
			render.N("footer", "Résulats vote..."),
		),
	)))
}

type merge struct {
	Register uint
	Votes    [driver2024legislative.OpinionLength]uint
}

func (m *merge) mergeBureauVote(bv *BureauVote) {
	m.Register += bv.Inscrit
	m.Votes[driver2024legislative.OpinionBlank] += bv.Blancs
	m.Votes[driver2024legislative.OpinionNull] += bv.Nuls
	m.Votes[driver2024legislative.OpinionAbstention] += bv.Abstentions
	m.mergeVotes(bv.Votes)
}

func (m *merge) mergeVotes(votes []driver2024legislative.Vote) {
	for _, v := range votes {
		m.Votes[v.Opinion] += uint(v.Voix)
	}
}
