package interp

import "fmt"

func Linear(t *table, j int, x float64) (y float64, err error) {
	if t.xx[j] == t.xx[j+1]{
		return t.yy[j],fmt.Errorf("interp: x values are not monotonic") // table is defective, but we can recover
	} else {
		y = t.yy[j] + ((x - t.xx[j])/(t.xx[j+1] - t.xx[j]))*(t.yy[j+1] - t.yy[j])
	}
	return y,nil
}
