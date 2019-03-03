
package main

import (
//	"fmt"
	"math"
)


func f_interp_k(delta float64) float64 {

	var a float64

	a = -1.0

	if math.Abs(delta) >= 0.0 && math.Abs(delta) <= 1.0 {
		return (a + 2.0)*math.Pow(math.Abs(delta), 3.0) - (a + 3.0)*math.Pow(math.Abs(delta), 2.0) + 1.0
	}

	if math.Abs(delta) > 1.0 && math.Abs(delta) < 2.0 {
		return a*math.Pow(math.Abs(delta), 3.0) + 5.0*math.Pow(math.Abs(delta), 2.0) + 8.0*a*math.Abs(delta) - 4.0*a
	}

/*
	if math.Abs(delta) >= 2.0 {
		return 0.0
	}
*/
	return 0.0
}

func f_interp_cubic(s []uint8, c float64, n int) float64 {
	var res float64 = 0.0
	var delta float64
	for k := 0; k < 4; k++ {
		delta = c - (float64)(n - 1 + k)
		res += (float64)(s[k]) * f_interp_k(delta)
	}
	return res
}


func f_interp_cubic_float64(s []float64, c float64, n int) float64 {
	var res float64 = 0.0
	var delta float64
	for k := 0; k < 4; k++ {
		delta = c - (float64)(n - 1 + k)
		res += s[k] * f_interp_k(delta)
	}
	return res
}

func f_interp_bicubic_float64(subdom st_domain, x float64, y float64) float64 {

	var tmp []float64 = nil
	var nx int = subdom.x_min + 1
	var ny int = subdom.y_min + 1
	tmp = make([]float64, 4)

	for k := 0; k < 4; k++ {
		data_row := subdom.get_data_row(k)
		tmp[k] = f_interp_cubic(data_row, x, nx)
	}

	return f_interp_cubic_float64(tmp, y, ny)

}


func f_interp_bicubic(subdom st_domain, x float64, y float64) uint8 {

	var tmp []float64 = nil
	var nx int = subdom.x_min + 1
	var ny int = subdom.y_min + 1
	tmp = make([]float64, 4)

	for k := 0; k < 4; k++ {
		data_row := subdom.get_data_row(k)
		tmp[k] = f_interp_cubic(data_row, x, nx)
	}

	return (uint8)(f_interp_cubic_float64(tmp, y, ny))

}



