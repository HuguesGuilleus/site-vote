package render

import (
	"fmt"
	"lfi/data-vote/common"

	"github.com/HuguesGuilleus/sniffle/tool/render"
)

var componentNav = render.Z
var componentFooter = render.Z

func renderBar(sum *common.Summary) render.Node {
	return render.N("div.bar", render.S2(sum.Result[:], "", func(_o int, voices uint) render.Node {
		o := common.Opinion(_o)
		if voices == 0 || o == common.OpinionAbstention {
			return render.Z
		}
		return render.Na("div.option",
			"data-opinion", o.String()).
			A("style", fmt.Sprintf("width:%d%%", percent(voices, sum.Register))).
			N()
	}))
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
