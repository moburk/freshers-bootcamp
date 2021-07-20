package main

import "fmt"

type Employee interface {
	calcSalary() int
}

type Fulltime struct{
	basePay int
}

type Freelancer struct {
	basePay, hours int
}
// USE COMPOSITION
func(emp Employee) {
	fmt.Println("Salary:", emp.calcSalary())
}

func main(){
	var emp1 Employee
	fte1 := Fulltime{500}
	emp1 = fte1
	fmt.Println("Full time employee salary:", emp1.calcSalary())
	contract1 := Fulltime{100}
	emp1 = contract1
	fmt.Println("Contractor salary:" , emp1.calcSalary())
	fmt.Printf("Enter hours worked by freelancer: ")
	var hrs int
	fmt.Scan(&hrs)
	freelance1 := Freelancer{10, hrs}
	emp1 = freelance1
	fmt.Println("Freelancer salary:" , emp1.calcSalary())
}

func (emp Fulltime) calcSalary() int {
	return emp.basePay*28
}

func (emp Freelancer) calcSalary() int {
	return emp.basePay*emp.hours
}