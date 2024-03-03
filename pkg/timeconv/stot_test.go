package timeconv

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type stotTest struct {
	name    string
	in      string
	out     time.Time
	isError bool
}

var stotTests = []stotTest{
	{"successful", "1586815793", time.Unix(1586815793, 0), false},
	{"invalid syntax", "abc123", time.Time{}, true},
	{"out of range", "9223372036854775808", time.Time{}, true},
}

func TestStot(t *testing.T) {
	for _, test := range stotTests {
		t.Run(test.name, func(t *testing.T) {
			out, err := Stot(test.in)

			require.Equal(t, test.out, out)

			if test.isError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func BenchmarkStot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Stot(stotTests[0].in)
	}
}
