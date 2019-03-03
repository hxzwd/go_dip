
package main


import (
	"fmt"
	"image"
	"image/png"
	"os"
	"math"
)


func append_for_2_pow(data *[]complex64) {

	var L int = len(*data)
	var p2 int
	int_p, _ := math.Modf(math.Sqrt((float64)(L)))
	p2 = (int)(int_p) + 1

	var L_new int = (int)(math.Pow(2.0, (float64)(p2)))
	for i := 0; i < L_new - L; i++ {
		*data = append(*data, 0)
	}

}

func run_test_1(file_name string) {

	s_img, err := form_st_img(file_name)

	if err != nil {
		panic(err.Error())
	}

	s_img.show_info()

	mat := s_img.make_matrix()

	fmt.Println("KSJDFSJKDF: \n ", len(mat))
	fmt.Println("ZKJDKJSD:  \n", len(mat[0]))

	err = s_img.save_as_image("out_imgs/orig.png")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(len(mat))

	var tmp_img st_8bpp_image

	tmp_img.from_matrix(mat)

	tmp_img.show_info()

	var new_gray *image.Gray = tmp_img.as_image()
	new_gray = new_gray
//	fmt.Println(new_gray)

	tmp_img = t_invert(s_img)


	err = tmp_img.save_as_image("out_imgs/out.png")

	if err != nil {
		panic(err.Error())
	}

	var dx int = 0
	var dy int = 1
	co_mat := calc_co_mat(s_img, dx, dy)
	co_mat = co_mat
//	fmt.Println(co_mat)


	tmp_img = t_transpose(tmp_img)
	err = tmp_img.save_as_image("out_imgs/out_t.png")
	if err != nil {
		panic(err.Error())
	}


	tmp_img = t_mirror_x(s_img)
	err = tmp_img.save_as_image("out_imgs/out_1.png")
	if err != nil {
		panic(err.Error())
	}


	tmp_img = t_mirror_y(s_img)
	err = tmp_img.save_as_image("out_imgs/out_2.png")
	if err != nil {
		panic(err.Error())
	}

	pdf := tmp_img.pdf()
	sum_of_pdf := 0.0
	for k := range pdf {
		sum_of_pdf += pdf[k]
	}
	fmt.Println(sum_of_pdf)

	fmt.Printf("mean: %f\n", s_img.m())
	fmt.Printf("var: %f\n", s_img.D())
	fmt.Printf("uni: %f\n", s_img.uniformity())

	s_img.save_hist("out_hist/1_hist.txt")

	err = s_img.save_co_mat("out_hist/2_co.txt", dx, dy)
	if err != nil {
		panic(err.Error())
	}

	s_img.save_co_mat("out_hist/3_co_1.txt", dy, dx)
	if err != nil {
		panic(err.Error())
	}

	test_dy := 4
	test_dx := 3
	s_img.save_co_mat("out_hist/4_co_2.txt", test_dx, test_dy)
	if err != nil {
		panic(err.Error())
	}

	s_img.save_co_mat("out_hist/5_co_3.txt", dx, dx)
	if err != nil {
		panic(err.Error())
	}

	var blend_imgs []string = []string{ "images/baboon.png", "images/Peppers.png", "images/boat.png"}
	test_blending(blend_imgs)
}

func main_old_1() {

	fmt.Printf("png_test golang version [main.go]\n\n")

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)



	file_name := "images/baboon.png"


	handle, err := os.Open(file_name)
	if err != nil {
		panic(err.Error())
	}

	defer handle.Close()

	pixels, err := get_pixels(handle)

	if err != nil {
		panic(err.Error())
	}

//	fmt.Println(pixels)
	pixels = pixels


/*
	byte_array, _, err := get_8bpp_array(os.Args[1])

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(byte_array)

*/

/*

	s_img, err := form_st_img(file_name)

	if err != nil {
		panic(err.Error())
	}

	s_img.show_info()

	mat := s_img.make_matrix()

	fmt.Println(len(mat))

	var tmp_img st_8bpp_image

	tmp_img.from_matrix(mat)

	tmp_img.show_info()

	var new_gray *image.Gray = tmp_img.as_image()
	new_gray = new_gray
//	fmt.Println(new_gray)

	tmp_img = t_invert(s_img)


	err = tmp_img.save_as_image("out.png")

	if err != nil {
		panic(err.Error())
	}

	var dx int = 0
	var dy int = 1
	co_mat := calc_co_mat(s_img, dx, dy)
	co_mat = co_mat
//	fmt.Println(co_mat)


	tmp_img = t_transpose(tmp_img)
	err = tmp_img.save_as_image("out_t.png")
	if err != nil {
		panic(err.Error())
	}


	tmp_img = t_mirror_x(s_img)
	err = tmp_img.save_as_image("out_1.png")
	if err != nil {
		panic(err.Error())
	}


	tmp_img = t_mirror_y(s_img)
	err = tmp_img.save_as_image("out_2.png")
	if err != nil {
		panic(err.Error())
	}

	pdf := tmp_img.pdf()
	sum_of_pdf := 0.0
	for k := range pdf {
		sum_of_pdf += pdf[k]
	}
	fmt.Println(sum_of_pdf)

	fmt.Printf("mean: %f\n", s_img.m())
	fmt.Printf("var: %f\n", s_img.D())
	fmt.Printf("uni: %f\n", s_img.uniformity())

	s_img.save_hist("hist/baboon_hist.txt")

	err = s_img.save_co_mat("hist/baboon_co.txt", dx, dy)
	if err != nil {
		panic(err.Error())
	}

	s_img.save_co_mat("hist/baboon_co_1.txt", dy, dx)
	if err != nil {
		panic(err.Error())
	}

	s_img.save_co_mat("hist/baboon_co_2.txt", dy, dy)
	if err != nil {
		panic(err.Error())
	}

	s_img.save_co_mat("hist/baboon_co_3.txt", dx, dx)
	if err != nil {
		panic(err.Error())
	}

	var blend_imgs []string = []string{ "images/baboon.png", "images/Peppers.png", "images/boat.png"}
	test_blending(blend_imgs)
*/


/*
	var c_data []complex64 = []complex64{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 + 1i, 12, 14 }
	fmt.Println(c_data)
	fmt.Println(len(c_data))
	append_for_2_pow(&c_data)
	fmt.Println(c_data)
	fmt.Println(len(c_data))

//	cmplx_bit_reverse(&c_data, len(c_data))

	fft(&c_data, len(c_data), 4, 1)

	fmt.Println(c_data)
/

	var test_file_name string = "images/Bagira.png"
	run_test_1(test_file_name)


	s_img, err := form_st_img(file_name)

	if err != nil {
		panic(err.Error())
	}

	s_img.show_info()

	s_img.show_all_stat()

/*

	tmp_img := t_rotate_img(s_img, 45.0)
	tmp_img.show_info()

	err = tmp_img.save_as_image("rot_out.png")
	if err != nil {
		panic(err.Error())
	}
*/

}

func rot_bagira() {
	file_name := "images/Bagira.png"
	out_file_pattern := "out_imgs/bagira_r_%d.png"
	var out_file string
	var r_img st_8bpp_image
	var angle float64

	s_img, err := form_st_img(file_name)

	if err != nil {
		panic(err.Error())
	}

	s_img.show_info()

	angle = 45.0
	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))

	r_img = t_rotate_img(s_img, angle)
	save_img(r_img, out_file)

	angle = -13.0
	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))

	r_img = t_rotate_img(s_img, angle)
	save_img(r_img, out_file)


}

func main() {


	file_name := "images/baboon.png"
	out_file_pattern := "out_imgs/test_baboon_r_%d.png"

	var r_img st_8bpp_image
	var out_file string
	var angle float64

	s_img, err := form_st_img(file_name)

	if err != nil {
		panic(err.Error())
	}

	s_img.show_info()

	angle = 45.0
	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))

	r_img = t_rotate_img(s_img, angle)

	err = r_img.save_as_image(out_file)

	if err != nil {
		panic(err.Error())
	}

	angle = -122.0
	r_img = t_rotate_img(s_img, angle)

	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))
	err = r_img.save_as_image(out_file)

	if err != nil {
		panic(err.Error())
	}


	angle = 13.0
	r_img = t_rotate_img(s_img, angle)

	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))
	err = r_img.save_as_image(out_file)

	if err != nil {
		panic(err.Error())
	}

	angle = 0.0
	r_img = t_rotate_img(s_img, angle)

	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))
	err = r_img.save_as_image(out_file)

	if err != nil {
		panic(err.Error())
	}

	angle = -90.0
	r_img = t_rotate_img(s_img, angle)

	out_file = fmt.Sprintf(out_file_pattern, (int)(angle))
	err = r_img.save_as_image(out_file)

	if err != nil {
		panic(err.Error())
	}



//	rot_bagira()

}

func save_img(img st_8bpp_image, filename string) {
	err := img.save_as_image(filename)
	if err != nil {
		panic(err.Error())
	}
}
