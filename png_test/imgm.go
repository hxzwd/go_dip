
package main

import (
	"math"
	"os"
	"fmt"
)


func (img *st_8bpp_image) hist() []uint32 {

	var hist []uint32

	hist = make([]uint32, (int)(math.Pow(2.0, (float64)(img.bpp))))

	for k := 0; k < img.size; k++ {
		hist[img.data[k]]++;
	}

	return hist

}


func (img *st_8bpp_image) pdf() []float64 {

	var hist []uint32

	hist = img.hist()
	var pdf []float64 = make([]float64, len(hist))


	for k := 0; k < len(hist); k++ {
		pdf[k] = (float64)(hist[k])/(float64)(img.size)
	}

	return pdf

}

func (img *st_8bpp_image) m() float64 {

	pdf := img.pdf()
	var m float64 = 0.0

	for k := range pdf {
		m += (float64)(k)*pdf[k]
	}

	return m

}

 func (img *st_8bpp_image) D() float64 {

	pdf := img.pdf()
	var D float64 = 0.0

	for k := range pdf {
		D += (float64)(k)*(float64)(k)*pdf[k]
	}

	return D

}

func (img *st_8bpp_image) uniformity() float64 {

	pdf := img.pdf()
	var r float64 = 0.0

	for k := range pdf {
		r += pdf[k]*pdf[k]
	}

	return r

}

func (img *st_8bpp_image) kurtosis() float64 {

	pdf := img.pdf()
	m := img.m()
	D := img.D()

	var res float64 = 0.0

	for k := range pdf {
		res += math.Pow(((float64)(k) - m), 4.0)*pdf[k]
	}

	return res/(D*D) - 3.0

}

func (img *st_8bpp_image) skewness() float64 {

	pdf := img.pdf()
	m := img.m()
	sigma := math.Sqrt(img.D())

	var res float64 = 0.0

	for k := range pdf {
		res += math.Pow(((float64)(k) - m), 3.0)*pdf[k]
	}

	return res/(sigma*sigma*sigma)

}

func (img *st_8bpp_image) entropy() float64 {

	pdf := img.pdf()

	var res float64 = 0.0
	var epsilon float64 = 0.000000001

	for k := range pdf {
		res += (pdf[k] + epsilon)*math.Log2(pdf[k] + epsilon)
	}

	return -res

}


func (img *st_8bpp_image) show_all_stat() {

	fmt.Printf("image: %s\n", img.name)
	fmt.Printf("[mean] I = %f\n", img.m())
	fmt.Printf("[var] sigma^2 = %f\n", img.D())
	fmt.Printf("[kurtosis] E = %f\n", img.kurtosis())
	fmt.Printf("[skewness] A = %f\n", img.skewness())
	fmt.Printf("[uniformity] E_n = %f\n", img.uniformity())
	fmt.Printf("[entropy] H_n = %f\n", img.entropy())

}





func (img *st_8bpp_image) save_hist(file_name string) error {

	var bytes_writed int = 0

	hist := img.hist()
	handle, err := os.OpenFile(file_name, os.O_WRONLY | os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer handle.Close()

	for k := range hist {
		bytes_w, err := handle.WriteString(fmt.Sprintf("%d %d\n", k, hist[k]))
		if err != nil {
			handle.Close()
			return err
		}
		bytes_writed += bytes_w
	}

	fmt.Printf("[%s] %d bytes writed in %s\n", "save_hist()", bytes_writed, file_name)

	return nil

}

func (img st_8bpp_image) save_co_mat(file_name string, dx int, dy int) error {

	var co_mat [][]uint32
	var bytes_writed int = 0

	co_mat = calc_co_mat(img, dx, dy)

	var n int = len(co_mat)
	var m int = len(co_mat[0])

//	fmt.Println(n)
//	fmt.Println(m)
//	fmt.Println(co_mat[2])

	handle, err := os.OpenFile(file_name, os.O_WRONLY | os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	defer handle.Close()

	for x := 0; x < n; x++ {
		var row_str string = "["
		for y := 0; y < m; y++ {
			if y == m - 1 {
				row_str += fmt.Sprintf("%d]\n", co_mat[x][y])
				continue
			}
			row_str += fmt.Sprintf("%d,", co_mat[x][y])
		}
//		fmt.Println(row_str)
		bytes_w, err := handle.WriteString(row_str)
		if err != nil {
			handle.Close()
			return err
		}
		bytes_writed += bytes_w
	}

	fmt.Printf("[%s] %d bytes writed in %s\n", "save_co_mat()", bytes_writed, file_name)

	return nil


}

