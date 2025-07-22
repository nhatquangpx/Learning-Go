package main 

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type PermanentEmployee struct {
	id int
	name string
	basic int
	pf int
}

type ContractEmployee struct {
	id int
	name string
	basic int
}

func (p PermanentEmployee) CalculateSalary() int {
	return p.basic + p.pf
}

func (c ContractEmployee) CalculateSalary() int {
	return c.basic
}

func totalExpense (e []SalaryCalculator) int {
	total := 0
	for _, v := range e {
		total = total + v.CalculateSalary()
	}
	return total
}

func main() {
	pemp1 := PermanentEmployee{id: 1, name: "John", basic: 5000, pf: 500}
	pemp2 := PermanentEmployee{id: 2, name: "Jane", basic: 6000, pf: 600}
	cemp1 := ContractEmployee{id: 3, name: "Doe", basic: 4000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	fmt.Printf("Total Expense: $%d", totalExpense(employees))
}