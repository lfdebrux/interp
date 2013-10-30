package interp

type table struct {
	xx []float64
	yy []float64
	n int
	ascnd bool // true if ascending order of table, false otherwise
}

type rawinterp func(t *table, j int, x float64) (y float64, err error)
type Interp func(x float64) (y float64, err error)

func New(xx, yy []float64, method rawinterp) Interp {
	if len(xx) != len(yy) {
		panic("interp: data x and y values of unequal length")
	}

	t := &table{
		xx,
		yy,
		len(xx),
		xx[len(xx)-1] >= xx[0],
	}

	return func(x float64) (y float64, err error) {
		j, err := t.locate(x)
		if err != nil {
			return 0,err
		}
		y, err = method(t,j,x)

		return y, err
	}
}

func (t *table) locate(x float64) (j int,err error) {

	var ju,jm,jl int

	jl = 0
	ju = t.n-1

	for ju-jl > 1 {
		jm = (ju+jl)/2 // find midpoint

		if x >= t.xx[jm] == t.ascnd {
			jl = jm
		} else {
			ju = jm
		}
	}

	return j, nil
}
