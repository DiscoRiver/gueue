package timeconv

import "time"

// Ttoi translates a time.Time value into a valid unix timestamp, represented as an int64.
func Ttoi(t time.Time) int64 {
	return t.Unix()
}
