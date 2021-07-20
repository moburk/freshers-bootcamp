package main

import "fmt"

const (
	fulltimebase = 500
	contractbase = 100
	freelancebase = 10
	workingdays = 28
)

type employeeSalary interface {
	calcSalary() int
}

type salaryCalculator struct {
	payPerUnitTime, timeWorked int
}

type fulltime struct{
	basePay int
	ISalaryCalculator salaryCalculator
}

type contractor struct{
	basePay int
	ISalaryCalculator salaryCalculator
}

type freelancer struct {
	basePay, hours int
	ISalaryCalculator salaryCalculator
}
// USE COMPOSITION

func main(){
	//var emp1 employeeSalary
	fte1 := fulltime{
		basePay: fulltimebase,
		ISalaryCalculator: salaryCalculator{
			payPerUnitTime: fulltimebase,
			timeWorked: workingdays,
		},
	}
	fmt.Println("Full time employee salary:", fte1.ISalaryCalculator.calcSalary())
	contract1 := contractor {
		basePay: contractbase,
		ISalaryCalculator: salaryCalculator{
			payPerUnitTime: contractbase,
			timeWorked: workingdays,
		},
	}
	fmt.Println("Contractor salary:" , contract1.ISalaryCalculator.calcSalary())
	fmt.Printf("Enter hours worked by freelancer: ")
	var hrs int
	fmt.Scan(&hrs)
	freelance1 := freelancer{
		basePay: freelancebase,
		hours: hrs,
		ISalaryCalculator: salaryCalculator{
			payPerUnitTime: freelancebase,
			timeWorked: hrs,
		},
	}
	fmt.Println("Freelancer salary:" , freelance1.ISalaryCalculator.calcSalary())
}

func (emp salaryCalculator) calcSalary() int {
	return emp.payPerUnitTime * emp.timeWorked
}