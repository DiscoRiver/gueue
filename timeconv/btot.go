package timeconv

import (
	"strconv"
	"time"
)

// Btot translates a unix time string, represented as bytes, into a time.Time value. Returns zero-value time.Time and error if parsing fails.
func Btot(b []byte) (time.Time, error) {
	i, err := strconv.ParseInt(string(b), base, intSize)
	if err != nil {
		return time.Time{}, err
	}
	tm := time.Unix(i, nsec)
	return tm, nil
}
