package render

import (
	"fmt"
	"lfi/data-vote/votation"
	"sniffle/tool"
	"sniffle/tool/render"
	"time"
)

var componentNav = render.Z
var componentFooter = render.Z

func RenderStation(t *tool.Tool, s *votation.Station) {
	p := fmt.Sprintf("/vote/%s/%s/%s.html", s.DépartementCode, s.City, s.CodeStation)
	t.WriteFile(p, render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=/vote/style.css>`),
			render.N("title", s.City, ": ", s.CodeStation),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../../..").N("vote"),
					render.Na("a.headerBlock", "href", "..").N(s.DépartementName),
					render.Na("a.headerBlock", "href", ".").N(s.City),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Bureau:", s.CodeStation),
				),
			),
			render.N("main",
				render.S(s.Votation, "", func(v votation.Votation) render.Node {
					sum := votation.NewSummary(&v.VotationResult)
					return render.N("",
						render.N("h1", "[", v.Date.Format(time.DateOnly), "] ", v.Name),
						render.N("div.bar", render.S2(sum.Result[:], "", func(o int, voices uint) render.Node {
							opi := votation.Opinion(o)
							if voices == 0 || opi == votation.OpinionAbstention {
								return render.Z
							}
							return render.Na("div.option",
								"data-opinion", opi.String()).
								A("style", fmt.Sprintf("width:%d%%", percent(voices, v.Register))).
								N()
						})),
						render.N("table",
							render.N("tr",
								render.N("th", "Voix"),
								render.N("th", "%"),
								render.N("th", "Nuance"),
								render.N("th", "Liste"),
							),
							render.S(v.Result, "", func(r votation.Result) render.Node {
								return render.N("tr",
									render.N("td.r", r.Result),
									render.N("td.r", percent(r.Result, v.Register), "%"),
									render.N("td",
										render.Na("div.copinion", "data-opinion", r.Opinion.String()),
										r.Party),
									render.N("td",
										"[", r.Position, "] ",
										r.FirstName, " ", r.LastName,
										" ", r.Gender.String(),
									),
								)
							}),
							render.N("tr",
								render.N("td.r", v.Blank),
								render.N("td.r", percent(v.Blank, v.Register), "%"),
								render.N("td",
									render.Na("div.copinion", "data-opinion", votation.OpinionBlank.String()),
									"Blanc"),
							),
							render.N("tr",
								render.N("td.r", v.Null),
								render.N("td.r", percent(v.Null, v.Register), "%"),
								render.N("td",
									render.Na("div.copinion", "data-opinion", votation.OpinionNull.String()),
									"Nul"),
							),
							render.N("tr",
								render.N("td.r", v.Abstention),
								render.N("td.r", percent(v.Abstention, v.Register), "%"),
								render.N("td",
									render.Na("div.copinion", "data-opinion", votation.OpinionAbstention.String()),
									"Abstention"),
							),
							render.N("tr",
								render.N("td.r", v.Register),
								render.N("td"),
								render.N("td", "Total"),
							),
						),
					)
				}),
			),

			componentFooter,
		),
	)))
}

func percent(part, total uint) uint {
	if part == 0 || total == 0 {
		return 0
	}
	p := part * 10 * 100 / total
	if p%10 > 4 {
		return p/10 + 1
	} else {
		return p / 10
	}
}
