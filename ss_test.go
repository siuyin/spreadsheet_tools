package ss

import "testing"

func TestCtoI(t *testing.T) {
	type d struct {
		i string
		o int
	}
	dat := []d{
		d{"B", 1},
		d{"AA", 26},
		d{"Ab", 27},
		d{"ba", 52},
		d{"ca", 78},
		d{"cb", 79},
		d{"ce", 82},
		d{"aaa", 702},
	}
	for k, v := range dat {

		if i := CtoI(v.i); i != v.o {
			t.Errorf("case: %d, %s, %d got %d", k, v.i, v.o, i)
		}
	}
}

func TestPow26(t *testing.T) {
	if i := pow26(0); i != 1 {
		t.Errorf("0: %d", i)
	}
	if i := pow26(1); i != 26 {
		t.Errorf("1: %d", i)
	}
	if i := pow26(2); i != 26*26 {
		t.Errorf("2: %d", i)
	}
}
