package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// first usage
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//==
//second usage
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//You cannot declare a method with a receiver whose type is defined in another package
// below is an illegal function
// func (f int) Abc() int {
// 	if f < 0 {
// 		return int(-f)
// 	}
// 	return int(f)

// }

//value reciever
func (v Vertex) NewAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//pointer recievers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Methods() {
	// v := Vertex{3, 4}
	// fmt.Println(v.Abs())
	// fmt.Println(Abs(v))

	// vertex := Vertex{48, 48}
	// fmt.Println(vertex.Abs())

	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())

	// v := Vertex{3, 4}
	// v.Scale(10)
	// fmt.Println(v.Abs())

	v := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.NewAbs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.NewAbs())
}

