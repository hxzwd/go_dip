
package main

import (
	"math"
	"fmt"
	"sort"
)


type st_pos struct {
	j int
	i int
	a uint8
}

func int_part(val float64) int {

	int_p, _ := math.Modf(val)

	return (int)(int_p)
}


func frac_part(val float64) float64 {

	_, frac_part := math.Modf(val)

	return frac_part
}


func int_coord(x float64, y float64) (int, int) {
	return int_part(x), int_part(y)
}


func calc_dist(x float64, y float64, i int, j int) float64 {

	var res float64 = 0.0
	var xx float64 = (float64)(i)
	var yy float64 = (float64)(j)

	res = math.Sqrt((x - xx)*(x - xx) + (y - yy)*(y - yy))

	return res

}

func get_16_n(x float64, y float64, img st_8bpp_image) ([][]st_pos, int, int) {

	var res [][]st_pos
	ii, jj := int_coord(x, y)
	mat := img.make_matrix()

	var xi, yi int
	var m_dist float64 = math.Sqrt(2.0)

	for i := -1; i <= 2; i++ {
		var row []st_pos
		for j := -1; j <= 2; j++ {
			if ii + i < 0 || ii + i >= img.h || jj + j < 0 || jj +j >= img.h {
				row = append(row, st_pos{ i + ii, j + jj, 0 })
			} else {
				row = append(row, st_pos{ i + ii, j + jj, mat[i + ii][j + jj] })
			}
		}
		res = append(res, row)

	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if res[i][j].i < 0 || res[i][j].j < 0 {
				continue
			}
			tmp_dist := calc_dist(x, y, res[i][j].i, res[i][j].j)
			if tmp_dist < m_dist {
				m_dist = tmp_dist
				xi = res[i][j].i
				yi = res[i][j].j
			}

		}
	}

	return res, xi, yi

}

func interp_calc_coeff(points []st_pos, y float64) float64 {

	var A [][]int = [][]int{{0, 6, 0, 0}, {-2, -3, 6, -1}, {3, -6, 3, 0}, {-1, 3, -3, 1}}
	var b[]float64
	var res float64 = 0.0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(points) ; j++ {
			b = append(b, (float64)(points[j].a) * (float64)(A[i][j]))
		}
	}

	for i := 0; i < len(b); i++ {
		res += b[i]*math.Pow(y, (float64)(i))
	}

	return res
}

func interp_calc_coeff_fast(points [][]st_pos, x float64, y float64) float64 {

	var A [][]int = [][]int{{0, 6, 0, 0}, {-2, -3, 6, -1}, {3, -6, 3, 0}, {-1, 3, -3, 1}}
//	var b []float64
//	var bb []float64
//	var bbb []float64
//	var bbbb []float64
	var res float64 = 0.0
	var res0, res1, res2, res3 float64
	var cc0, cc1, cc2, cc3 float64
	res0 = 0.0
	res1 = 0.0
	res2 = 0.0
	res3 = 0.0

	for i := 0; i < len(A); i++ {
		cc0 = 0.0
		cc1 = 0.0
		cc2 = 0.0
		cc3 = 0.0
		for j := 0; j < len(points) ; j++ {
			cc0 += ((float64)(points[0][j].a) * (float64)(A[i][j]))
			cc1 += ((float64)(points[1][j].a) * (float64)(A[i][j]))
			cc2 += ((float64)(points[2][j].a) * (float64)(A[i][j]))
			cc3 += ((float64)(points[3][j].a) * (float64)(A[i][j]))
		}
		res0 += cc0 * math.Pow(y, (float64)(i))
		res1 += cc1 * math.Pow(y, (float64)(i))
		res2 += cc2 * math.Pow(y, (float64)(i))
		res3 += cc3 * math.Pow(y, (float64)(i))
	}

	cc0 = 0.0
	cc1 = 0.0
	cc2 = 0.0
	cc3 = 0.0
	cc0 = (float64)(A[0][1]) + (float64)(A[0][1]) * res1 + (float64)(A[0][2]) * res2 * res2 +  (float64)(A[0][3]) * res3 * res3 * res3
	cc1 = (float64)(A[1][1]) + (float64)(A[1][1]) * res1 + (float64)(A[1][2]) * res2 * res2 +  (float64)(A[1][3]) * res3 * res3 * res3
	cc2 = (float64)(A[2][1]) + (float64)(A[2][1]) * res1 + (float64)(A[2][2]) * res2 * res2 +  (float64)(A[2][3]) * res3 * res3 * res3
	cc3 = (float64)(A[3][1]) + (float64)(A[3][1]) * res1 + (float64)(A[3][2]) * res2 * res2 +  (float64)(A[3][3]) * res3 * res3 * res3


	res = cc0  + cc1 * x + cc2 * x * x + cc3 * x * x * x

	return res
}


func interp_by_new_y(points []float64, x float64) float64 {

	var A [][]int = [][]int{{0, 6, 0, 0}, {-2, -3, 6, -1}, {3, -6, 3, 0}, {-1, 3, -3, 1}}
	var b[]float64
	var res float64 = 0.0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(points) ; j++ {
			b = append(b, points[j] * (float64)(A[i][j]))
		}
	}

	for i := 0; i < len(b); i++ {
		res += b[i]*math.Pow(x, (float64)(i))
	}

	return res

}

func get_new_z(x float64, y float64, img st_8bpp_image) (int, int, uint8) {

	points, xi, yi := get_16_n(x, y, img)

//	fmt.Println(points)

	var new_values []float64
	new_values = make([]float64, 4)

	for i := 0; i < 4; i++ {
		new_values[i] = interp_calc_coeff(points[i], frac_part(y))
	}

	z := (uint8)(interp_by_new_y(new_values, frac_part(x)))


//	fmt.Println(xi, yi, z)
	return xi, yi, z

}


func get_new_z_fast(x float64, y float64, img st_8bpp_image) (int, int, uint8) {

	points, xi, yi := get_16_n(x, y, img)

//	fmt.Println(points)


	z := (uint8)(interp_calc_coeff_fast(points, frac_part(x), frac_part(y)))

//	fmt.Println(xi, yi, z)
	return xi, yi, z

}


func t_rotate_img_backup(img st_8bpp_image, angle float64) st_8bpp_image {

	fmt.Println("Rotate: ", angle)

	var res st_8bpp_image
	res.copy_from(img)

	var x, y float64
	mat := img.make_matrix()

	angle = angle*math.Pi/180.0

	for old_x := 0; old_x < img.h; old_x++ {
		fmt.Println("old_x = ", old_x)
		for old_y := 0; old_y < img.w; old_y++ {
			xx := (float64)(old_x)
			yy := (float64)(old_y)
			x = xx*math.Cos(angle) + yy*math.Sin(angle)
			y = yy*math.Cos(angle) - xx*math.Sin(angle)
			if x >= 0.0 && y >= 0.0 && x <= (float64)(img.h - 1) && y <= (float64)(img.w - 1) {
				xi, yi, z := get_new_z_fast(x, y, img)
				xi = xi
				yi = yi
				mat[old_x][old_y] = z
//				fmt.Println(old_x, old_y, x, y, z);
			} else {
				mat[old_x][old_y] = 0
//				fmt.Println(old_x, old_y, x, y, 0);
			}
		}
	}

//	xi, yi, z := get_new_z(x, y, img)

//	fmt.Printf("\n\n%d %d\n", xi, yi)
//	fmt.Printf("z = %d\n", z)

	res.from_matrix(mat)
	return res

}




func t_rotate_img_old_1(img st_8bpp_image, angle float64) st_8bpp_image {

	fmt.Println("Rotate: ", angle)

	var res st_8bpp_image
	res.copy_from(img)

	var x, y float64
	mat := img.make_matrix()

	angle = angle*math.Pi/180.0

	for old_x := 0; old_x < img.h; old_x++ {
		fmt.Println("old_x = ", old_x)
		for old_y := 0; old_y < img.w; old_y++ {
			xx := (float64)(old_x)
			yy := (float64)(old_y)
			x = xx*math.Cos(angle) + yy*math.Sin(angle)
			y = yy*math.Cos(angle) - xx*math.Sin(angle)
			if x >= 0.0 && y >= 0.0 && x <= (float64)(img.h - 1) && y <= (float64)(img.w - 1) {
				xi, yi, z := get_new_z(x, y, img)
				xi = xi
				yi = yi
				mat[old_x][old_y] = z
//				fmt.Println(old_x, old_y, x, y, z);
			} else {
				mat[old_x][old_y] = 0
//				fmt.Println(old_x, old_y, x, y, 0);
			}
		}
	}

//	xi, yi, z := get_new_z(x, y, img)

//	fmt.Printf("\n\n%d %d\n", xi, yi)
//	fmt.Printf("z = %d\n", z)

	res.from_matrix(mat)
	return res
}

func rot_point(x0 int, y0 int, angle float64) (float64, float64) {

	xx := (float64)(x0)
	yy := (float64)(y0)
	x := xx*math.Cos(angle) + yy*math.Sin(angle)
	y := yy*math.Cos(angle) - xx*math.Sin(angle)
	return x, y
}

func MinMax(a []float64) (float64, float64) {
	var b []float64 = nil
	b = make([]float64, len(a))
	copy(b[:], a)
	sort.Float64s(b)
	return b[0], b[len(b) - 1]
}

func MinMaxInt(a []float64) (int, int) {
	x, y := MinMax(a)
	return (int)(x), (int)(y)
}

func new_bound(img st_8bpp_image, angle float64) []int {

	x_max := img.w
	y_max := img.h

	x0, y0 := rot_point(0, 0, angle)
	x1, y1 := rot_point(x_max, 0, angle)
	x2, y2 := rot_point(x_max, y_max, angle)
	x3, y3 := rot_point(0, y_max, angle)

	xb_min, xb_max := MinMaxInt([]float64{x0, x1, x2, x3})
	yb_min, yb_max := MinMaxInt([]float64{y0, y1, y2, y3})
/*
	xb_max := math.Max(x0, x1, x2, x3)
	xb_min := math.Min(x0, x1, x2, x3)
	yb_max := math.Max(y0, y1, y2, y3)
	yb_min := math.Min(y0, y1, y2, y3)
*/

/*
	fmt.Println(x0, y0)
	fmt.Println(x1, y1)
	fmt.Println(x2, y2)
	fmt.Println(x3, y3)
	fmt.Println("Bound:")
	fmt.Println(xb_min, yb_min)
	fmt.Println(xb_max, yb_min)
	fmt.Println(xb_max, yb_max)
	fmt.Println(xb_min, yb_max)
*/
	res := []int{xb_min, xb_max, yb_min, yb_max}

	return res

}

func t_rotate_img_old_2_test(img st_8bpp_image, angle float64) st_8bpp_image {

	fmt.Println("Rotate: ", angle)

	var res st_8bpp_image
	var dom st_domain
//	mat := img.make_matrix()

	angle = angle*math.Pi/180.0

	bound_coords := new_bound(img, angle)
	dom.init(bound_coords)
	fmt.Println(dom.get_params())

	for j := dom.y_min; j <= dom.y_max; j++ {
		for i := dom.x_min; i <= dom.x_max; i++ {
			x_r, y_r := rot_point(i, j, -angle)
			if img.in_image_f(x_r, y_r) {
				dom.set_value(i, j, 255)
			} else {
				dom.set_value(i, j, 0)
			}
		}
	}


//TEST BLOCK
	x_test, y_test := 250, 100
	x_test_r, y_test_r := rot_point(x_test, y_test, -angle)
	fmt.Printf("\n\nTEST [-angle = %f rad]:\n%d %d\n%f %f\n", -angle, x_test, y_test, x_test_r, y_test_r)
	fmt.Println("In image: ", img.in_image_f(x_test_r, y_test_r))

	var dom_test st_domain
	var test_file_name string = "out_imgs/test_img.png"
	dom_test.from_matrix(img.make_matrix(), []int{0, 511, 0, 511}, 2)
	dom_test.save_as_image(test_file_name)

	var subdom_test st_domain
	subdom_test = dom_test.get_rect_subdomain(x_test_r, y_test_r, 4)
	fmt.Println(subdom_test.get_params())
	fmt.Println(subdom_test.data)
	subdom_test = dom_test.get_rect_subdomain(10.1, 0.1, 4)
	fmt.Println(subdom_test.get_params())
	fmt.Println(subdom_test.data)
	fmt.Println("subdom_test.get_data_row:\n", "row = 1;\t", subdom_test.get_data_row(1), "\nrow = 3;\t", subdom_test.get_data_row(3))
	fmt.Println("subdom_test.get_domain_row:\n", "row = -1;\t", subdom_test.get_domain_row(-1), "\nrow = 1;\t", subdom_test.get_domain_row(1))

	var interp_result_float64 float64
	var interp_result uint8
	interp_result_float64 = f_interp_bicubic_float64(subdom_test, 10.1, 0.1)
	interp_result = f_interp_bicubic(subdom_test, 10.1, 0.1)
	fmt.Printf("interp_result = %d[%f]\n", interp_result, interp_result_float64)

	res.from_matrix(dom.data)
	return res

}


func t_rotate_img(img st_8bpp_image, angle float64) st_8bpp_image {

	fmt.Println("Rotate: ", angle)

	var res st_8bpp_image
	var dom st_domain
	var dom_src st_domain
	var subdomain st_domain
	var border_size int
	var subdomain_size int

	border_size = 2
	subdomain_size = 4

	angle = angle*math.Pi/180.0

	bound_coords := new_bound(img, angle)
	fmt.Println("Bound coords:\n", bound_coords)

	dom.init(bound_coords)
	fmt.Println("Destination domain:\n", dom.get_params())
	dom_src.from_matrix(img.make_matrix(), []int{0, img.w - 1, 0, img.h - 1}, border_size)
	fmt.Println("Source domain:\n", dom_src.get_params())

	for j := dom.y_min; j < dom.y_max; j++ {
		for i := dom.x_min; i < dom.x_max; i++ {
			x_r, y_r := rot_point(i, j, -angle)
			if img.in_image_f(x_r, y_r) {
				subdomain = dom_src.get_rect_subdomain(x_r, y_r, subdomain_size)
				dom.set_value(i, j, f_interp_bicubic(subdomain, x_r, y_r))
			} else {
				//fmt.Printf("i = %d\tj = %d\n", i, j)
				dom.set_value(i, j, 0)
			}
		}
	}

	res.from_matrix(dom.data)
	return res

}
