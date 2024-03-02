package timeconv

import (
	"testing"
	"time"
)

type btotTest struct {
	in  []byte
	out time.Time
	err error
}

var btotTests = []btotTest{
	{[]byte("1586815793"), time.Unix(1586815793, 0), nil},
	{[]byte("1570390523"), time.Unix(1570390523, 0), nil},
	{[]byte("1546094442"), time.Unix(1546094442, 0), nil},
}

func TestBtot(t *testing.T) {
	for _, test := range btotTests {
		out, err := Btot(test.in)
		if test.out != out || !(test.err == err) {
			t.Errorf("Btot(%q) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

func BenchmarkBtot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Btot(btotTests[0].in)
	}
}
