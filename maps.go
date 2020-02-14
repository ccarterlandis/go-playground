package main

import (
	// "fmt"
	"strings"
	"golang.org/x/tour/wc"
)

func Maps() {
	// m = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }
	// fmt.Println(m["Bell Labs"])
	// fmt.Println(m)
	// mutateMaps()
	wc.Test(WordCount)
}

// It should return a map of the counts of each “word” in the string s. 
//The wc.Test function runs a test suite against the provided function and prints success or failure. 
func WordCount(s string) map[string]int {
	results := map[string]int{}
	fields := strings.Fields(s)
	ok := false
	for _, value := range fields {
		_, ok = results[value]
		if ok == false {
			results[value] = 1
		} else {
			results[value] = results[value] + 1
		}
	}
	return results
}

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex

// func mutateMaps() {
// 	m := make(map[string]int)

// 	m["Answer"] = 42
// 	fmt.Println("The value:", m["Answer"])

// 	m["Answer"] = 48
// 	fmt.Println("The value:", m["Answer"])

// 	delete(m, "Answer")
// 	fmt.Println("The value:", m["Answer"])

// 	v, ok := m["Answer"]
// 	fmt.Println("The value:", v, "Present?", ok)
// }


