package main

import "fmt"

type Matrix struct {
	rows, cols int
	elements [][]int
}

func (m Matrix) init() {
	for i := range m.elements {
		m.elements[i] = make([]int, m.cols)
	}
}

func (m *Matrix) setValues() {
	for i:=0; i<m.rows; i++ {
		for j:=0; j<m.cols; j++{
			m.elements[i][j] = i*(m.rows-1) + j
		}
	}
}

func (m Matrix) getRows() int {
	return m.rows
}

func (m Matrix) getColumns() int {
	return m.cols
}

func (m Matrix) setElement(i, j, val int) {
	if i < 0 || j < 0 || i >= m.rows || j >= m.cols {
		fmt.Println("Index out of bounds")
		return
	}
	m.elements[i][j] = val
}

func (m Matrix) addMatrices (n Matrix) Matrix {
	ans := Matrix{m.rows, m.cols, make([][]int, m.rows)}
	ans.init()
	for i := range m.elements {
		for j := range m.elements[i] {
			ans.elements[i][j] = m.elements[i][j] + n.elements[i][j]
		}
	}
	return ans
}

func main(){
	fmt.Println("Enter rows and columns: ")
	var rows, cols int
	fmt.Scan(&rows)
	fmt.Scan(&cols)
	matrix1 := Matrix{ rows, cols, make([][]int, rows)}
	matrix1.init()
	fmt.Println(matrix1)
	fmt.Println(matrix1.getRows())
	fmt.Println(matrix1.getColumns())
	matrix1.setElement(1, 1, 5)
	fmt.Println(matrix1.elements)
	matrix2 := Matrix{ rows, cols, make([][]int, rows)}
	matrix2.init()
	matrix2.setValues()
	fmt.Println(matrix2)
	fmt.Println(matrix1.addMatrices(matrix2))
}
