package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Matrix struct {
	Rows int `json:"rows"` // uncapitalized fields are ignored by json.Marshal
	Cols int `json:"cols"`
	Elements [][]int `json:"elements"`
}

func (m *Matrix) init() {
	for i := range m.Elements {
		m.Elements[i] = make([]int, m.Cols)
	}
}

func (m *Matrix) setValues() {
	for i:=0; i<m.Rows; i++ {
		for j:=0; j<m.Cols; j++{
			m.Elements[i][j] = i*(m.Rows-1) + j
		}
	}
}

func (m Matrix) getRows() int {
	return m.Rows
}

func (m Matrix) getColumns() int {
	return m.Cols
}

func (m *Matrix) setElement(i, j, val int) error {
	if i < 0 || j < 0 || i >= m.Rows || j >= m.Cols {
		return errors.New("index out of bounds")
	}
	m.Elements[i][j] = val
	return nil
}

func (m Matrix) addMatrices (n Matrix) Matrix {
	ans := Matrix{m.Rows, m.Cols, make([][]int, m.Rows)}
	ans.init()
	for i := range m.Elements {
		for j := range m.Elements[i] {
			ans.Elements[i][j] = m.Elements[i][j] + n.Elements[i][j]
		}
	}
	return ans
}

func (m Matrix) printJson() {
	//fmt.Printf("%+v", m)
	json1, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json1))
}

func main(){
	fmt.Println("Enter rows and columns: ")
	var rows, cols int
	fmt.Scan(&rows)
	fmt.Scan(&cols)
	matrix1 := Matrix{ rows, cols, make([][]int, rows)}
	matrix1.init()
	fmt.Println(matrix1)
	fmt.Println("Rows =", matrix1.getRows())
	fmt.Println("Columns =", matrix1.getColumns())
	err := matrix1.setElement(2, 1, 5)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(matrix1.Elements)
	matrix2 := Matrix{ rows, cols, make([][]int, rows)}
	matrix2.init()
	matrix2.setValues()
	fmt.Println(matrix2)
	fmt.Println(matrix1.addMatrices(matrix2))
	matrix2.printJson()
}
