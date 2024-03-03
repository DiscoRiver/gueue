package timeconv

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type btotTest struct {
	name    string
	in      []byte
	out     time.Time
	isError bool
}

var btotTests = []btotTest{
	{"successful", []byte("1586815793"), time.Unix(1586815793, 0), false},
	{"syntax error", []byte("abc123"), time.Time{}, true},
	{"out of range", []byte("9223372036854775808"), time.Time{}, true},
}

func TestBtot(t *testing.T) {
	for _, test := range btotTests {
		t.Run(test.name, func(t *testing.T) {
			out, err := Btot(test.in)

			require.Equal(t, test.out, out)

			if test.isError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func BenchmarkBtot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Btot(btotTests[0].in)
	}
}
