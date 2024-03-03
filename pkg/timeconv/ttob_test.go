package timeconv

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type ttobTest struct {
	name string
	in   time.Time
	out  []byte
}

var ttobTests = []ttobTest{
	{"successful", time.Unix(1586815793, 0), []byte("1586815793")},
	{"successful zero time", time.Unix(0, 0), []byte("0")},
}

func TestTtob(t *testing.T) {
	for _, test := range ttobTests {
		t.Run(test.name, func(t *testing.T) {
			out := Ttob(test.in)

			require.Equal(t, test.out, out)
		})
	}
}

func BenchmarkTtob(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ttob(ttobTests[0].in)
	}
}
