package main

import (
	"lfi/data-vote/driver2024legislative"
	"sniffle/tool"
)

func main() {
	t := tool.New(tool.CLI(nil))
	driver2024legislative.Parse(t)
}
