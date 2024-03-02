package timeconv

import (
	"fmt"
	"time"
)

// Ttob translates a time.Time value into a valid unix timestamp, represented in bytes.
func Ttob(t time.Time) []byte {
	return []byte(fmt.Sprintf("%d", t.Unix()))
}
