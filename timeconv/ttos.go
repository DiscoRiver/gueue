package timeconv

import (
	"fmt"
	"time"
)

// Ttos translates a time.Time value into a valid unix timestamp, represented as a string.
func Ttos(t time.Time) string {
	return fmt.Sprintf("%d", t.Unix())
}
