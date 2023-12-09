package tfl

import "fmt"

// HTTPError can be returned to represent a failed HTTP call.
type HTTPError struct {
	Status string
	Body   string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Body)
}
