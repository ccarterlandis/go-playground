package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for index := range picture {
		picture[index] = make([]uint8, dx)
		//fmt.Println(index, row)
	}
	for y_index := range picture {
		for x_index := range picture[y_index] {
			picture[y_index][x_index] = uint8((x_index ^ y_index))
		}
	}
	return picture
}

func PictureMain() {
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from not being done yet.", r)
        }
    }()
	pic.Show(Pic)
}