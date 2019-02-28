package main

import (
	"fmt"
	"image"
//	"image/color"
	"image/png"
	"os"
)

type st_8bpp_image struct {

	data []uint8
	w int
	h int
	size int
	name string
	bpp int

}


func (img *st_8bpp_image) show_info() {

	fmt.Printf("[image name] name = %s\n", img.name)
	fmt.Printf("[image bpp] bpp = %d\n", img.bpp)
	fmt.Printf("[image width] w = %d\n", img.w)
	fmt.Printf("[image height] h = %d\n", img.h)
	fmt.Printf("[image total size] size = w * h = %d * %d = %d\n", img.w, img.h, img.size)
}

func (img *st_8bpp_image) make_matrix() [][]uint8 {

	row_size := img.w
	columns_num := img.h

	k := 0
	var mat [][]uint8

	for x := 0; x < columns_num; x++ {
		var row []uint8
		for y := 0; y < row_size; y++ {
			row = append(row, img.data[k])
			k++
		}
		mat = append(mat, row)
	}

	return mat

}

func (img *st_8bpp_image) from_matrix(mat [][]uint8) {

	img.h = len(mat)
	img.w = len(mat[0])
	img.size = img.h * img.w

	var data []uint8


	for x := 0; x < img.h; x++ {
		for y := 0; y < img.w; y++ {
			data = append(data, mat[x][y])
		}
	}

	img.data = data
	img.name = "empty"
	img.bpp = 8

}

func (img *st_8bpp_image) as_image() *image.Gray  {

	raw_data := ([]byte)(img.data)
//	fmt.Println(raw_data)
	res := image.NewGray(image.Rect(0, 0, img.w, img.h))

	res.Pix = raw_data

	return res

}

func (img *st_8bpp_image) copy_from(src st_8bpp_image)  {
	img.w = src.w
	img.h = src.h
	img.size = src.size
	img.name = src.name
	var tmp_data []uint8
	for i := 0; i < src.size; i++ {
		tmp_data  = append(tmp_data, src.data[i])
	}
	img.data = tmp_data
	img.bpp = src.bpp

}

func (img *st_8bpp_image) save_as_image(file_name string) error  {

	tmp_img := img.as_image()

	handle, err := os.OpenFile(file_name, os.O_WRONLY | os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer handle.Close()

	png.Encode(handle, tmp_img)

	return nil

}

func (img *st_8bpp_image) at_fn(pos int) float64  {

	var res float64
	res = (float64)(img.data[pos])/255.0

	return res
}

func (img *st_8bpp_image) set_fn(val float64, pos int)  {

	var res uint8
	res = (uint8)((float64)(val * 255.0))

	img.data[pos] = res
}
