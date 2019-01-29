package structs

import "time"

type Employee struct {
	EmployeeId int
	Name string
	Address string
	DOB time.Time
	Position string
	Salary int
	ManagerId int
}

var adarsh = Employee{1, "Adarsh", "Adarsh Addr", time.Now(), "Backend Engg.", 1, 2}