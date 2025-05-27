package csvtool

import (
	"fmt"
	"strconv"
)

func ParseUint(s string) uint {
	if s == "-1" {
		return 0
	}
	u, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("ParseUint(%q): %s", s, err))
	}
	return uint(u)
}
