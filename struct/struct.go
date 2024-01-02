package main

import "fmt"

// จัดการโครงสร้างด้วย Struct
type employee struct {
	employeeID string
	employeeName string
	employeePhone string
}

func main() {
	//struct by slice
	employeeList := []employee{}
	employee1 := employee{
		employeeID: "101",
		employeeName: "Tle Tle",
		employeePhone: "0900000000",
	}
	employee2 := employee{
		employeeID: "102",
		employeeName: "Tee Tee",
		employeePhone: "0900000001",
	}
	employee3 := employee{
		employeeID: "103",
		employeeName: "To To",
		employeePhone: "0900000002",
	}

	employeeList = append(employeeList, employee1, employee2, employee3)
	// employeeList = append(employeeList, employee2)
	// employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)


	// array
	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeID: "101",
	// 	employeeName: "Tle Tle",
	// 	employeePhone: "0900000000",
	// }
	// employeeList[1] = employee{
	// 	employeeID: "102",
	// 	employeeName: "Tee Tee",
	// 	employeePhone: "0900000001",
	// }
	// employeeList[2] = employee{
	// 	employeeID: "103",
	// 	employeeName: "To To",
	// 	employeePhone: "0900000002",
	// }
	// fmt.Println("employee = ", employeeList)
	
}