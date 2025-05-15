package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// CalculateArea, UserInfo
	// Structs, interfaces, enums

	// snake_case
	// user_id, first_name

	// UPPERCASE
	// Use case is constants

	// mixedCase
	// htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
