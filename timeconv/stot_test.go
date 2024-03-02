package timeconv

import (
	"testing"
	"time"
)

type stotTest struct {
	in  string
	out time.Time
	err error
}

var stotTests = []stotTest{
	{"1586815793", time.Unix(1586815793, 0), nil},
	{"1570390523", time.Unix(1570390523, 0), nil},
	{"1546094442", time.Unix(1546094442, 0), nil},
}

func TestStot(t *testing.T) {
	for _, test := range stotTests {
		out, err := Stot(test.in)
		if test.out != out || !(test.err == err) {
			t.Errorf("Stot(%q) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

func BenchmarkStot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Stot(stotTests[0].in)
	}
}
