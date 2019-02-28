
package main

import (
//	"fmt"
	"math"
//	"math/cmplx"
)


func swap_in_array(data *[]complex64, pos1 int, pos2 int) {

	tmp := (*data)[pos1]
	(*data)[pos1] = (*data)[pos2]
	(*data)[pos2] = tmp

}

func cmplx_bit_reverse(data *[]complex64, size int) {

	var middle int
	var reverse_size int
	var i, j int

	middle = size / 2
	reverse_size = size - 1
	j = 0

	for i = 0; i < reverse_size; i++ {
		if i < j {
			swap_in_array(data, i, j)
		}
		var k int = middle
		for ; k <= j;  {
			j -= k
			k /= 2
		}
		j += k
	}

}


func fft(data *[]complex64, size int, k int, dir int) {

	var p_left, p_right int
	p_left = 1
	p_right = 1

	cmplx_bit_reverse(data, size)

	for stage := 1; stage <= k; stage++ {
		p_left = p_right
		p_right *= 2
		var c_rot complex64 = 1.0 + 0i
		var c_arg float64 = math.Pi / (float64)(p_left)
		var f_dir float64 = (float64)(dir)
		var w_factor complex64 = complex((float32)(math.Cos(c_arg)), (float32)(-math.Sin(c_arg)*f_dir))
		for butter_pos := 0; butter_pos < p_left; butter_pos++ {
			for top_node := butter_pos; top_node < size; top_node += p_right {
				var bot_node int = top_node + p_left
				var temp complex64 = (*data)[bot_node] * c_rot
				(*data)[bot_node] = (*data)[top_node] - temp
				(*data)[top_node] += temp
			}
			c_rot *= w_factor
		}
	}

}
