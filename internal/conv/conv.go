package conv

import (
	"fmt"
	"strconv"
	"strings"
)

func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}

func StringSliceToString(s []string) string {
	return strings.Join(s, ",")
}
