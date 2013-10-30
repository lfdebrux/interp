package interp

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func Poly(t *table, mm int, j int, x float64) (y float64, err error) {
	var i, m, ns int
	var den, dif, dift, ho, hp, w float64

	j -= (mm - 2)/2 // change index so that x is centred in m indices

	xa, ya := t.xx[j:], t.yy[j:]

	c, d := make([]float64,mm), make([]float64,mm)

	dif = abs(x-xa[0])
	for i=0;i<mm;i++ {
		if (dift=abs(x-xa[i])) < dif {
			ns = i
			dif = dift
		}
		c[i] = ya[i]
		d[i] = ya[i]
	}
	y = ya[ns--]
	for m=1;m<mm;m++ {
		for i=0;i<mm-m;i++ {
			ho=xa[i]-x
			hp=xa[i+m]-x
			w=c[i+1]-d[i]
			if (den=ho-hp) == 0.0 {
				panic("poly_interp error")
			}
			den=w/den
			d[i]=hp*den
			c[i]=ho*den
		}
		y += (dy=(2*(ns+1) < (mm-m) ? c[ns+1] : d[ns--]))
	}
	return y,nil
}
