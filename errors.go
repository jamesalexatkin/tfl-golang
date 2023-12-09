package tfl

import "fmt"

type HTTPError struct {
	Status string
	Body   string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Body)
}
