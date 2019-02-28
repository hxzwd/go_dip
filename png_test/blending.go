
package main

import (
	"fmt"
	"math"
)

func b_norm(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	return C_s
}

func b_mul(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		res.set_fn(C_b.at_fn(k) * C_s.at_fn(k), k)
	}
	return res
}


func b_screen(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		c := 1.0 - (1.0 - a)*(1.0 - b)
		res.set_fn(c, k)
	}
	return res
}


func b_darken(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		c := math.Min(a, b)
		res.set_fn(c, k)
	}
	return res
}


func b_lighten(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		c := math.Max(a, b)
		res.set_fn(c, k)
	}
	return res
}


func b_difference(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		c := math.Abs(a - b)
		res.set_fn(c, k)
	}
	return res
}


func b_color_dodge(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		if C_s.data[k] == 255 {
			res.set_fn(1.0, k)
		} else {
			c := math.Min(1.0, a/(1.0 - b))
			res.set_fn(c, k)
		}
	}
	return res
}


func b_soft_light(C_b st_8bpp_image, C_s st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		if C_s.data[k] == 0 {
			res.set_fn(0.0, k)
		} else {
			c := 1.0 - math.Min(1.0, (1.0 - a)/b)
			res.set_fn(c, k)
		}
	}
	return res
}

func c_img_alpha(C_b st_8bpp_image, C_s st_8bpp_image, alpha st_8bpp_image) st_8bpp_image {
	var res st_8bpp_image
	res.copy_from(C_b)
	for k := range res.data {
		a := C_b.at_fn(k)
		b := C_s.at_fn(k)
		alpha := alpha.at_fn(k)
		c := (1.0 - alpha)*a + alpha*b
		res.set_fn(c, k)

	}
	return res
}

func test_blending(file_names []string) error {

	var imgs []st_8bpp_image
	imgs = make([]st_8bpp_image, len(file_names))

	for k := range file_names {
		tmp_img, err := form_st_img(file_names[k])
		imgs[k] = tmp_img
		if err != nil {
			return err
		}
	}

	fmt.Printf("\n\n[test_blending]\n")
	fmt.Println("Compose 1-st and 2-nd images with 3-nd as alpha channel")

	C_b := imgs[0]
	C_s := imgs[1]
	alpha := imgs[2]
	var res_img st_8bpp_image

	res_img = c_img_alpha(C_b, C_s, alpha)

	res_img.save_as_image("c_alpha.png")

	return nil

}

