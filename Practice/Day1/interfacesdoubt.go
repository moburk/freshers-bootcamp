package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	t := &T{"Hello"}
	i = t
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	x := i // here is x a variable of type interface??? no I think
	describe(x)
	x = t
	describe(x)
	var j I
	y := j
	x = j
	describe(x) // here both x and y have type <nil> and value <nil>
	describe(y)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

