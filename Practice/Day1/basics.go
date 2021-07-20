package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	//"golang.org/x/tour/pic"
)

func swap(x, y string) (string, string) {
	return y, x
}
func fsqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func swapPtr(x, y *string) {
	tmp := *x
	*x = *y
	*y = tmp
}

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
	}
	for i := range picture {
		for j := range picture[i] {
			picture[i][j] = uint8(i)+uint8(j)
		}
	}
	return picture
}

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	words := strings.Fields(s)
	for i := range words {
		ans[words[i]]++
	}
	return ans
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}


func main() {
	fmt.Println("Random number is ", rand.Intn(10))
	fmt.Printf("Square root of 7: %f\n", math.Sqrt(7) )
	x, y := "world", "hello"
	fmt.Println(x, y)
	fmt.Println(swap(x, y)) // everything is passed by value unless slices or pointers are used
	fmt.Println(fsqrt(9))
	swapPtr(&x, &y)
	fmt.Println(x, y)
	//pic.Show(Pic)
	fmt.Println(WordCount("Able was I ere I saw Elba"))
	fn := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fn())
	}
}