package data

var employees []Employee // på tisdag gör om till databas  (SQLite + MySQL )

func GetAllEmployees() []Employee {
	return employees
}

func CreateNewEmployee(newEmployee Employee) {
	employees = append(employees, newEmployee)
}

func GetEmployee(id int) *Employee {
	for _, employee := range employees {
		if employee.Id == id {
			return &employee
		}
	}
	return nil
}

func Init() {

	employees = append(employees, Employee{Id: 1, Age: 52, Namn: "Stefan", City: "Test"})
	employees = append(employees, Employee{Id: 2, Age: 16, Namn: "Oliver", City: "Test"})
	employees = append(employees, Employee{Id: 3, Age: 22, Namn: "Josefine", City: "Test"})
}
