
package main

import (
	"fmt"
)


type st_domain struct {

	x_max int
	x_min int
	y_max int
	y_min int
	data [][]uint8
	x_len int
	y_len int
}

func (dom *st_domain) init(coords []int) {
	dom.x_min = coords[0]
	dom.x_max = coords[1]
	dom.y_min = coords[2]
	dom.y_max = coords[3]
	dom.x_len = dom.x_max - dom.x_min + 1
	dom.y_len = dom.y_max - dom.y_min + 1
	for i := 0; i < dom.y_len; i++ {
		var row []uint8
		for j := 0; j < dom.x_len; j++ {
			row = append(row, 0)
		}
		dom.data = append(dom.data, row)
	}
}

func (dom *st_domain) get_params() string {
	return fmt.Sprintf("x: %d %d\ny: %d %d\nlen: %d %d", dom.x_min, dom.x_max, dom.y_min, dom.y_max, dom.x_len, dom.y_len)
}

func (dom *st_domain) get_value(x int, y int) uint8 {
	if x < dom.x_min || x > dom.x_max || y < dom.y_min || y > dom.y_max {
		panic(fmt.Sprintf("st_domain: get_value error[%d %d]\nparams: \n%s\n", x, y, dom.get_params()))
	}
	xx := x - dom.x_min
	yy := y - dom.y_min
	if xx < 0 || yy < 0 || xx >= dom.x_len || yy >= dom.y_len {
		panic(fmt.Sprintf("st_domain: get_value error[%d %d]\nxx = %d\tyy = %d\nparams: \n%s\n", x, y, xx, yy, dom.get_params()))
	}
//	return dom.data[xx][yy]
	return dom.data[yy][xx]
}



func (dom *st_domain) set_value(x int, y int, value uint8) uint8 {
	if x < dom.x_min || x > dom.x_max || y < dom.y_min || y > dom.y_max {
		panic(fmt.Sprintf("st_domain: get_value error[%d %d]\nparams: \n%s\n", x, y, dom.get_params()))
	}
	xx := x - dom.x_min
	yy := y - dom.y_min
	if xx < 0 || yy < 0 || xx >= dom.x_len || yy >= dom.y_len {
		panic(fmt.Sprintf("st_domain: get_value error[%d %d]\nxx = %d\tyy = %d\nparams: \n%s\n", x, y, xx, yy, dom.get_params()))
	}
//	dom.data[xx][yy] = value
	dom.data[yy][xx] = value
	return value
}

func (dom *st_domain) get_rect_subdomain(x_c float64, y_c float64, size int) st_domain {

//	var res [][]uint8
	var subdom st_domain
	x0 := int_part(x_c) - size/2 + 1
	y0 := int_part(y_c) - size/2 + 1

	x1 := x0 + size - 1
	y1 := y0 + size - 1
/*
	for j := y0; j <= y1; j++ {
		var row []uint8
		for i := x0; i <= x1; i++ {
			row = append(row, dom.data[i][j])
		}
		res = append(res, row)
	}
*/
	subdom.init([]int{x0, x1, y0, y1})

//	fmt.Println("In get_rect_subdomain:\n", "sub params:\n", subdom.get_params(), "\ndom params:\n", dom.get_params())

	for j := y0; j <= y1; j++ {
		for i := x0; i <= x1; i++ {
//			fmt.Printf("i = %d\tj = %d\n", i, j)
//			subdom.set_value(i, j, dom.data[i][j])
//			subdom.set_value(i, j, dom.get_value(i, j))
			subdom.set_value(i, j, dom.get_value(i, j))
		}
	}

	return subdom

}

func (dom *st_domain) from_matrix(data [][]uint8, coords []int, border int) {
	var new_coords []int = nil
	new_coords = make([]int, len(coords))
	copy(new_coords[:], coords)
	new_coords[0] -= border
	new_coords[2] -= border
	new_coords[1] += border
	new_coords[3] += border
	dom.init(new_coords)

//	fmt.Println(dom.get_params())
/*
	fmt.Println("len data: ", len(data))
	fmt.Println("len data[0]: ", len(data[0]))
	fmt.Println("len dom.data: ", len(dom.data))
	fmt.Println("len dom.data[0]: ", len(dom.data[0]))
	fmt.Println("dom.y_len - border: ", dom.y_len - border)
	fmt.Println("dom.x_len - border: ", dom.x_len - border)
*/
	for j := border; j < dom.y_len - border; j++ {
		for i := border; i < dom.x_len - border; i++ {
			dom.data[j][i] = data[j - border][i - border]
//			dom.data[i][j] = data[i - border][j - border]
//			fmt.Printf("i = %d\tj = %d\ni - border = %d\tj - border = %d\n", i, j, i - border, j - border)
		}
	}

}

func (dom *st_domain) save_as_image(file_name string) {
	var img st_8bpp_image
	img.from_matrix(dom.data)
	err := img.save_as_image(file_name)
	if err != nil {
		panic(fmt.Sprintf("st_domain [save_as_image]\nfile_name = %s\nparams:\n%s", file_name, dom.get_params()))
	}
	fmt.Printf("From st_domain save in image: %s\n", file_name)
}

func (dom *st_domain) get_data_row(j_pos int) []uint8 {
	return dom.data[:][j_pos]
}

func (dom *st_domain) get_domain_row(y_pos int) []uint8 {
	if y_pos < dom.y_min || y_pos > dom.y_max {
		panic(fmt.Sprintf("st_domain: get_domain_row error[%d]\nparams: \n%s\n", y_pos, dom.get_params()))
	}
	yy := y_pos - dom.y_min
	return dom.data[:][yy]
}

