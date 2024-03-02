package timeconv

import (
	"strconv"
	"time"
)

// Stot translates a unix time string into a Go time.Time value. Returns zero-value time.Time and error if parsing fails.
func Stot(s string) (time.Time, error) {
	i, err := strconv.ParseInt(s, base, intSize)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(i, nsec), nil
}
