package main

import (
	"fmt"
	"math"
	"math/rand"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Random number is ", rand.Intn(10))
	fmt.Printf("Square root of 7: %f\n", math.Sqrt(7) )
	x, y := "world", "hello"
	fmt.Println(x, y)
	fmt.Println(swap(x, y)) // everything is passed by value unless slices or pointers are used
}