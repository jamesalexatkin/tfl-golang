package conv

import "strings"

// StringSliceToString converts a string slice to a string.
func StringSliceToString(s []string) string {
	return strings.Join(s, ",")
}
