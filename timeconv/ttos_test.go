package timeconv

import (
	"testing"
	"time"
)

type ttosTest struct {
	in  time.Time
	out string
}

var ttosTests = []ttosTest{
	{time.Unix(1586815793, 0), "1586815793"},
	{time.Unix(1570390523, 0), "1570390523"},
	{time.Unix(1546094442, 0), "1546094442"},
}

func TestTtos(t *testing.T) {
	for _, test := range ttosTests {
		out := Ttos(test.in)
		if test.out != out {
			t.Errorf("Ttob(%q) = %v want %v",
				test.in, out, test.out)
		}
	}
}

func BenchmarkTtos(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ttos(ttosTests[0].in)
	}
}
