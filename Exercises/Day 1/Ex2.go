package main

import "fmt"

func checkVar (s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func preorder(s string) {
	for i, ch := range s {

	}
}

func main(){
	var exp string = "a+b-c"
	fmt.Println(exp)
}
