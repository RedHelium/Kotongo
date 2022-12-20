package extensions

import (
	"strconv"
	"strings"
)

func Concat(args ...string) string {

	var str strings.Builder

	for i := range args {
		str.WriteString(args[i])
	}

	return str.String()
}

func FloatToHex(val float32) string {
	return strconv.FormatInt(int64(val), 16)
}
