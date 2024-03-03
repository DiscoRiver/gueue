package timeconv

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type ttosTest struct {
	name string
	in   time.Time
	out  string
}

var ttosTests = []ttosTest{
	{"successful", time.Unix(1586815793, 0), "1586815793"},
	{"successful zero time", time.Unix(0, 0), "0"},
}

func TestTtos(t *testing.T) {
	for _, test := range ttosTests {
		t.Run(test.name, func(t *testing.T) {
			out := Ttos(test.in)
			
			require.Equal(t, test.out, out)
		})

	}
}

func BenchmarkTtos(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ttos(ttosTests[0].in)
	}
}
