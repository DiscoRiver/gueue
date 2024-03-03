package timeconv

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type ttoiTest struct {
	name string
	in   time.Time
	out  int64
}

var ttoiTests = []ttoiTest{
	{"successful", time.Unix(1586815793, 0), 1586815793},
	{"successful zero time", time.Unix(0, 0), 0},
}

func TestTtoi(t *testing.T) {
	for _, test := range ttoiTests {
		t.Run(test.name, func(t *testing.T) {
			out := Ttoi(test.in)

			require.Equal(t, test.out, out)
		})
	}
}

func BenchmarkTtoi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ttoi(ttoiTests[0].in)
	}
}
