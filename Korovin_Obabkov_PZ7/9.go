package main

import (
	"fmt"
)

func main() {
	Department1 := Department{"IT", []Employee{}}
	E1 := Employee{"Андрей", "Сисадмин", 90000}
	E2 := Employee{"Дмитрий", "Программист", 40000}
	AddEmployee(&Department1, E1)
	AddEmployee(&Department1, E2)
	fmt.Println("сотрудники до увольнения;")
	fmt.Print(GetEmployeesByPosition(Department1, "Программист"))
	fmt.Print(GetEmployeesByPosition(Department1, "Сисадмин"))
	fmt.Printf("\nсумма зарплат в IT департаменте: %v", GetSumSalary(Department1))
	FireEmployee(&Department1, "Терентий")
	fmt.Println("\nсисадмины после увольнения:")
	fmt.Println(GetEmployeesByPosition(Department1, "Сисадмин"))
}

type Employee struct {
	Name     string
	Position string
	Salary   float64
}
type Department struct {
	Name      string
	Employees []Employee
}

func AddEmployee(department *Department, employee Employee) {
	department.Employees = append(department.Employees, employee)
}
func FireEmployee(department *Department, name string) {
	for i, employee := range department.Employees {
		if employee.Name == name {
			department.Employees = append(department.Employees[:i], department.Employees[i+1:]...)
			break
		}
	}
}
func GetEmployeesByPosition(department Department, position string) []Employee {
	employees := []Employee{}
	for _, employee := range department.Employees {
		if employee.Position == position {
			employees = append(employees, employee)
		}
	}
	return employees
}
func GetSumSalary(department Department) float64 {
	sum := 0.0
	for _, employee := range department.Employees {
		sum += employee.Salary
	}
	return sum
}
