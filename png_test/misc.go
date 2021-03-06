
package main


import (
	"image"
	"io"
	"os"
)


type st_pixel struct {
	R int
	G int
	B int
	A int
}

type st_image_info struct {

	width uint32
	height uint32
	size uint32

}


func rgba_to_pixel(r uint32, g uint32, b uint32, a uint32) st_pixel {
	return st_pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func get_pixels(file io.Reader) ([][]st_pixel, error) {

	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()

	width := bounds.Max.X
	height := bounds.Max.Y

	var pixels_array [][]st_pixel

	for y := 0; y < height; y++ {
		var row []st_pixel
		for x := 0; x < width; x++ {
			row = append(row, rgba_to_pixel(img.At(x, y).RGBA()))
		}
		pixels_array = append(pixels_array, row)
	}

	return pixels_array, nil

}

func get_8bpp_array(file_name string) ([]uint8, st_image_info, error) {

	var img_byte []uint8
	var img_info st_image_info

	img_info.width = 0
	img_info.height = 0
	img_info.size = 0

	handle, err := os.Open(file_name)

	if err != nil {
		return nil, img_info, err
	}

	defer handle.Close()

	img, _, err := image.Decode(handle)

	if err != nil {
		return nil, img_info, err
	}

	bounds := img.Bounds()

	width := (uint32)(bounds.Max.X)
	height := (uint32)(bounds.Max.Y)
	total_size := width * height

	var x, y uint32

	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			tmp_pixel := rgba_to_pixel(img.At(int(x), int(y)).RGBA())
			img_byte = append(img_byte, (uint8)(tmp_pixel.R))
		}
	}

	return img_byte, st_image_info{ width, height, total_size }, nil

}


func form_st_img(file_name string) (st_8bpp_image, error) {

	var img st_8bpp_image
	var img_info st_image_info
	var data []uint8

	img.data = nil
	img.w = 0
	img.h = 0
	img.size = 0
	img.bpp = 8
	img.name = file_name

	data, img_info, err := get_8bpp_array(file_name)

	if err != nil {
		return img, err
	}

	img.data = data
	img.w = (int)(img_info.width)
	img.h = (int)(img_info.height)
	img.size = (int)(img_info.size)

	return img, nil


}



func t_invert(img st_8bpp_image) st_8bpp_image {

	var res st_8bpp_image
	res.copy_from(img)

	for i := 0; i < res.size; i++ {
		res.data[i] = 0xFF - img.data[i]
	}

	return res

}

func get_co_mat_(n int, m int) [][]uint32 {
	var res [][]uint32
	for x := 0; x < n; x++ {
		var row []uint32
		for y := 0; y < m; y++ {
			row = append(row, 0)
		}
		res = append(res, row)
	}
	return res
}



func make_zero_mat(n int, m int) [][]uint8 {
	var res [][]uint8
	for x := 0; x < n; x++ {
		var row []uint8
		for y := 0; y < m; y++ {
			row = append(row, 0)
		}
		res = append(res, row)
	}
	return res
}

func calc_co_mat(img st_8bpp_image, dx int, dy int) [][]uint32 {

	mat := img.make_matrix()
	var co_mat [][]uint32
	n := 256
	m := 256

	co_mat = get_co_mat_(n, m)

	width := img.w
	height := img.h

	for x := 0; x < height; x++ {

		for y := 0; y < width; y++ {
			pix1 := (int)(mat[x][y])
			if x + dx < height && y + dy < width {
				pix2 := (int)(mat[x + dx][y + dy])
				co_mat[pix1][pix2]++
			}
		}

	}

	return co_mat

}


func t_transpose(img st_8bpp_image) st_8bpp_image {

	var res st_8bpp_image
	res.w = img.h
	res.h = img.w
	res.size = img.size
	res.name = img.name + ":transpose"
	res.bpp = img.bpp

//	k := 0
	mat := img.make_matrix()
	mat2 := make_zero_mat(img.w, img.h)

	for x := 0; x < img.h; x++ {
		for y := 0; y < img.w; y++ {
			mat2[y][x] = mat[x][y]
		}
	}

	res.from_matrix(mat2)
	return res

}

func t_mirror_x(img st_8bpp_image) st_8bpp_image {

	var res st_8bpp_image
	res.w = img.h
	res.h = img.w
	res.size = img.size
	res.name = img.name + ":transpose"
	res.bpp = img.bpp

//	k := 0
	mat := img.make_matrix()
	mat2 := make_zero_mat(img.h, img.w)

	for x := 0; x < img.h; x++ {
		for y := 0; y < img.w; y++ {
			mat2[img.h - x - 1][y] = mat[x][y]
		}
	}

	res.from_matrix(mat2)
	return res

}

func t_mirror_y(img st_8bpp_image) st_8bpp_image {

	var res st_8bpp_image
	res.w = img.h
	res.h = img.w
	res.size = img.size
	res.name = img.name + ":transpose"
	res.bpp = img.bpp

//	k := 0
	mat := img.make_matrix()
	mat2 := make_zero_mat(img.h, img.w)

	for x := 0; x < img.h; x++ {
		for y := 0; y < img.w; y++ {
			mat2[x][img.w - y - 1] = mat[x][y]
		}
	}

	res.from_matrix(mat2)
	return res

}

