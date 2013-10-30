package interp

import "testing"

func TestLinear(t *testing.T) {
	xx := []float64 {0, 2, 4, 6,  8, 10}
	yy := []float64 {2, 4, 6, 8, 10, 12}

	f := New(xx,yy,Linear)

	test := map[float64]float64 {1:3, 3:5, 7:9, 9:11}

	for in,out := range test {
		if x,_ := f(in); x != out {
			t.Errorf("linear interpolation f(%v) = %v, should be %v", in, x, out)
		}
	}
}