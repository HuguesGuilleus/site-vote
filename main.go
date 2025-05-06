package main

import (
	_ "embed"
	"lfi/data-vote/driver2024legislative"
	"lfi/data-vote/render"
	"sniffle/tool"
)

//go:embed style.css
var style []byte

func main() {
	t := tool.New(tool.CLI(nil))
	bureaux := driver2024legislative.Parse(t)

	t.WriteFile("/vote/style.css", style)

	for _, bv := range bureaux {
		if bv.DépartementCode != "10" {
			continue
		}
		render.RenderBV(t, bv)
	}
}
