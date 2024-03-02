package timeconv

import (
	"reflect"
	"testing"
	"time"
)

type ttobTest struct {
	in  time.Time
	out []byte
}

var ttobTests = []ttobTest{
	{time.Unix(1586815793, 0), []byte("1586815793")},
	{time.Unix(1570390523, 0), []byte("1570390523")},
	{time.Unix(1546094442, 0), []byte("1546094442")},
}

func TestTtob(t *testing.T) {
	for _, test := range ttobTests {
		out := Ttob(test.in)
		if !reflect.DeepEqual(test.out, out) {
			t.Errorf("Ttob(%q) = %v want %v",
				test.in, out, test.out)
		}
	}
}

func BenchmarkTtob(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ttob(ttobTests[0].in)
	}
}
