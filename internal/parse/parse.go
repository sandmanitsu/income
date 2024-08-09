package parse

import (
	"strconv"
	"strings"
)

func Parse(flag string) (string, float64) {
	flagParams := (strings.Split(flag, "="))

	var flagValue float64
	if len(flagParams) == 2 {
		flagValue, _ = strconv.ParseFloat(
			strings.TrimSpace(flagParams[1]),
			64,
		)
	}

	return flagParams[0], flagValue
}
