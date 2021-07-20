package main

import (
	"fmt"
	"log"
)

type node struct{
	val byte
	lNode, rNode *node
}

func checkOperator (s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

func buildTree (s string) *node{
	if s == "" {
		return nil
	}
	if len(s) == 1 {
		if checkOperator(s) {
			log.Fatalln("Incorrect expression format")
		}
		root := node{s[0], nil, nil}
		return &root
	}
	if !checkOperator(string(s[1])) {
		log.Fatalln("Incorrect expression format")
	}
	left := &node{s[0], nil, nil}
	right := buildTree(s[2:])
	root := node{s[1], left, right}
	return &root
}

func preorder(root *node) {
	if root != nil {
		fmt.Printf("%c ", root.val)
		preorder(root.lNode)
		preorder(root.rNode)
	}
}

func postorder(root *node) {
	if root != nil {
		postorder(root.lNode)
		postorder(root.rNode)
		fmt.Printf("%c ", root.val)
	}
}

func main(){
	var exp string
	fmt.Print("Enter expression: ")
	fmt.Scan(&exp)
	var root *node
	root = buildTree(exp)
	fmt.Println("Preorder Traversal:")
	preorder(root)
	fmt.Println("\nPostorder Traversal:")
	postorder(root)
}
