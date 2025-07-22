package main

import "fmt"

type Employee struct {
	id int
	name string
	salary int
}

func (e Employee) DisplayInfo() {
	fmt.Printf ("Employee ID: %d, Name: %s, Salary: $%d\n", e.id, e.name, e.salary)
}

func (e Employee) UpdateSalaryValue(newSalary int) {
	e.salary = newSalary
	fmt.Printf("Updated Salary for %s: $%d\n", e.name, e.salary)
}

func (e *Employee) UpdateSalaryPointer(newSalary int) {
	e.salary = newSalary
	fmt.Printf("Updated Salary for %s: $%d\n", e.name, e.salary)
}

func main() { 
	emp1 := Employee{id: 1, name: "Alice", salary: 50000}
	emp2 := Employee{id: 2, name: "Bob", salary: 60000}

	emp1.DisplayInfo()
	emp2.DisplayInfo()

	emp1.UpdateSalaryValue(55000)
	emp2.UpdateSalaryPointer(65000)

	emp1.DisplayInfo()
	emp2.DisplayInfo()
}