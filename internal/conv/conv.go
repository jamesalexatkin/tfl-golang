package conv

import "strings"

func StringSliceToString(s []string) string {
	return strings.Join(s, ",")
}
