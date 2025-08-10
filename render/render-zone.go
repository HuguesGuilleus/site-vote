package render

import (
	"lfi/data-vote/common"

	"github.com/HuguesGuilleus/sniffle/tool"
	"github.com/HuguesGuilleus/sniffle/tool/render"
)

func RenderFrance(t *tool.Tool, z *common.Zone) {
	t.WriteFile("/index.html", render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=style.css>`),
			render.H(`<link rel=icon href=favicon.ico>`),
			render.N("title", z.Departement.String()),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "France"),
				),
			),
			renderZoneMain(z),
			componentFooter,
		),
	)))
}

func RenderDepartement(t *tool.Tool, z *common.Zone) {
	p := z.Departement.Code() + "/index.html"
	t.WriteFile(p, render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=../style.css>`),
			render.H(`<link rel=icon href=../favicon.ico>`),
			render.N("title", z.Departement.String()),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../index.html").N("France"),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Département: ", z.Departement),
				),
			),
			renderZoneMain(z),
			componentFooter,
		),
	)))
}

func RenderLegislativeDisctrict(t *tool.Tool, z *common.Zone) {
	p := z.Departement.Code() + "/" + z.District + ".html"
	t.WriteFile(p, render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=../style.css>`),
			render.H(`<link rel=icon href=../favicon.ico>`),
			render.N("title", "Circonscrption ", z.District, " de ", z.Departement),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../index.html").N("France"),
					render.Na("a.headerBlock", "href", "./index.html").N(z.Departement),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Circonscription: ", z.District),
				),
			),
			renderZoneMain(z),
			componentFooter,
		),
	)))
}

func RenderCity(t *tool.Tool, z *common.Zone) {
	p := z.Departement.Code() + "/" + z.City + "/index.html"
	t.WriteFile(p, render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=../../style.css>`),
			render.H(`<link rel=icon href=../../favicon.ico>`),
			render.N("title", z.City),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../../index.html").N("France"),
					render.Na("a.headerBlock", "href", "../index.html").N(z.Departement),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Ville: ", z.City),
				),
			),
			renderZoneMain(z),
			componentFooter,
		),
	)))
}

func RenderStation(t *tool.Tool, z *common.Zone) {
	p := z.Departement.Code() + "/" + z.City + "/" + z.StationID + ".html"
	t.WriteFile(p, render.Merge(render.Na("html", "lang", "fr").N(
		render.N("head",
			render.H(`<meta charset=utf-8>`),
			render.H(`<meta name=viewport content="width=device-width,initial-scale=1">`),
			render.H(`<link rel=stylesheet href=../../style.css>`),
			render.H(`<link rel=icon href=../../favicon.ico>`),
			render.N("title", z.City, ": ", z.StationID),
		),
		render.N("body",
			componentNav,
			render.N("header",
				render.N("div.headerRow",
					render.Na("a.headerBlock", "href", "../../index.html").N("France"),
					render.Na("a.headerBlock", "href", "../index.html").N(z.Departement),
					render.Na("a.headerBlock", "href", "./index.html").N(z.City),
				),
				render.N("div.headerRow",
					render.N("div.headerBlock.main", "Bureau: ", z.StationID),
				),
			),
			renderZoneMain(z),
			componentFooter,
		),
	)))
}

func renderZoneMain(z *common.Zone) render.Node {
	return render.N("main",
		render.N("div.summary", render.S(z.Vote, "", func(v common.Vote) render.Node {
			return render.N("",
				render.Na("a", "href", "#"+v.ID).N(v.ID),
				renderBar(&v.Summary),
			)
		})),

		render.If(len(z.Parents) != 0, func() render.Node {
			return render.N("ul.sub", render.S(z.Parents, "", func(p string) render.Node {
				return render.N("li", render.Na("a", "href", "../"+p+".html").N(p))
			}))
		}),
		render.If(len(z.Same) != 0, func() render.Node {
			return render.N("ul.sub", render.S(z.Same, "", func(sub string) render.Node {
				return render.N("li", render.Na("a", "href", sub+".html").N(sub))
			}))
		}),
		render.If(len(z.Sub) != 0, func() render.Node {
			return render.N("ul.sub", render.S(z.Sub, "", func(sub string) render.Node {
				return render.N("li", render.Na("a", "href", sub+"/index.html").N(sub))
			}))
		}),

		render.S(z.Vote, "", func(v common.Vote) render.Node {
			return render.N("",
				render.Na("h1", "id", v.ID).N(v.Name),
				render.IfElse(len(v.SubSummary) != 0, func() render.Node {
					return render.N("div.summary",
						render.N("a.r", "*"), renderBar(&v.Summary),
						render.S(v.SubSummary, "", func(s common.SubSummary) render.Node {
							return render.N("",
								render.Na("a.r", "href", s.Group+".html").N(s.Group),
								renderBar(&s.Summary),
							)
						}),
					)
				}, func() render.Node {
					return renderBar(&v.Summary)
				}),
				renderVoteTable(&v),
			)
		}),
	)
}

func renderVoteTable(v *common.Vote) render.Node {
	return render.N("table",
		render.N("tr",
			render.N("th", "Voix"),
			render.N("th", "%"),
			render.N("th", "Nuance"),
			render.N("th", "Liste"),
		),
		render.S(v.Option, "", func(r common.Option) render.Node {
			if r.Result == 0 {
				return render.Z
			}
			return render.N("tr",
				render.N("td.r.wnowrap", r.Result),
				render.N("td.r.wnowrap", percent(r.Result, v.Register), "%"),
				render.N("td.wnowrap",
					render.Na("div.copinion", "data-opinion", r.Opinion.String()),
					r.Party),
				render.N("td",
					"[", r.Position, "] ",
					r.Name,
					r.Gender.String(),
				),
			)
		}),
		render.If(len(v.Option) == 0, func() render.Node {
			return render.N("", render.S2(v.Summary.Result[:common.OpinionBlank], "", func(_o int, voices uint) render.Node {
				if voices == 0 {
					return render.Z
				}
				o := common.Opinion(_o)
				return render.N("tr",
					render.N("td.r.wnowrap", voices),
					render.N("td.r.wnowrap", percent(voices, v.Register), "%"),
					render.N("td.wnowrap",
						render.Na("div.copinion", "data-opinion", o.String()),
						"["+o.String()+"]"),
				)
			}))
		}),
		render.N("tr",
			render.N("td.r.wnowrap", v.Blank),
			render.N("td.r.wnowrap", percent(v.Blank, v.Register), "%"),
			render.N("td.wnowrap",
				render.Na("div.copinion", "data-opinion", common.OpinionBlank.String()),
				"Blanc"),
		),
		render.N("tr",
			render.N("td.r.wnowrap", v.Null),
			render.N("td.r.wnowrap", percent(v.Null, v.Register), "%"),
			render.N("td.wnowrap",
				render.Na("div.copinion", "data-opinion", common.OpinionNull.String()),
				"Nul"),
		),
		render.N("tr",
			render.N("td.r.wnowrap", v.Abstention),
			render.N("td.r.wnowrap", percent(v.Abstention, v.Register), "%"),
			render.N("td.wnowrap",
				render.Na("div.copinion", "data-opinion", common.OpinionAbstention.String()),
				"Abstention"),
		),
		render.N("tr",
			render.N("td.r.wnowrap", v.Register),
			render.N("td"),
			render.N("td", "Total"),
		),
	)
}
