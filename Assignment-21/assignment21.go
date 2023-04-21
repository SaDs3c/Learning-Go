package main

import "fmt"

const (
	numberOfEmployee int = 2
)

type employee struct {
	first      string
	last       string
	employeeID int
	nin        string
	title      string
}

type employees struct {
	employee1 employee
	employee2 employee
}

func main() {
	emp := employees{}

	for i := 0; i < numberOfEmployee; i++ {
		if i == 0 {
			fmt.Printf("Enter the employee's first name: ")
			fmt.Scanln(&emp.employee1.first)

			fmt.Printf("Enter the employee's last name: ")
			fmt.Scanln(&emp.employee1.last)

			fmt.Printf("Enter the employee's ID number: ")
			fmt.Scanln(&emp.employee1.employeeID)

			fmt.Printf("Enter the last four digits of the employee's NIN: ")
			fmt.Scanln(&emp.employee1.nin)

			fmt.Printf("Enter the employee's job title (do not include the word 'Engineer'):")
			fmt.Scanln(&emp.employee1.title)

		}

		if i == 1 {

			fmt.Printf("Enter the employee's first name: ")
			fmt.Scanln(&emp.employee2.first)

			fmt.Printf("Enter the employee's last name: ")
			fmt.Scanln(&emp.employee2.last)

			fmt.Printf("Enter the employee's ID number: ")
			fmt.Scanln(&emp.employee2.employeeID)

			fmt.Printf("Enter the last four digits of the employee's NIN: ")
			fmt.Scanln(&emp.employee2.nin)

			fmt.Printf("Enter the employee's job title (do not include the word 'Engineer'):")
			fmt.Scanln(&emp.employee2.title)

		}

	}

	for i := 0; i < numberOfEmployee; i++ {

		if i == 0 {
			fmt.Println("Employee information for", emp.employee1.first, emp.employee1.last)
			fmt.Println("ID:", emp.employee1.employeeID)
			fmt.Println("NIN:", emp.employee1.nin)
			fmt.Println("Title:", emp.employee1.title, "Engineer\n")
		}
		if i == 1 {
			fmt.Println("Employee information for", emp.employee2.first, emp.employee2.last)
			fmt.Println("ID:", emp.employee2.employeeID)
			fmt.Println("NIN", emp.employee2.nin)
			fmt.Println("Title:", emp.employee2.title, "Engineer")
		}

	}
}
