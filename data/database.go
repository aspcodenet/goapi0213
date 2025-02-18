package data

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetAllEmployees() []Employee {
	// HÄMTA ALLA  EMPLOYEES FRÅN EN DATABAS
	// SELECT * FROM EMPLOYEES
	// DATA FRÅN DATABASEN -> MAPPAS OM TILL OBJEKT (struktar)
	// relationsdatabas data där mappas till GO-objekt
	// ORM - Object Relational Mapping
	// ORM kpdbibliotek som för just ORM
	// Ef Core (Entity Framework) C# , JPA Java, Sequalize Node.js
	// GORM - ORM för GO
	var employees []Employee
	db.Find(&employees) // det som Find - SELECT * FROM Employees
	return employees
}

func UpdateEmployee(employee Employee) bool {
	var dbEmployee Employee
	err := db.First(&dbEmployee, employee.Id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	dbEmployee.Age = employee.Age
	dbEmployee.Namn = employee.Namn
	dbEmployee.City = employee.City
	db.Save(&employee)
	return true
}

func CreateNewEmployee(employee Employee) *Employee {
	db.Create(&employee) // INSERT INTO EMPLOYEES (AGE, NAMN, CITY) VALUES (employee.Age, employee.Namn, employee.City)
	return &employee
}
func GetEmployee(id int) *Employee { // GetEpmployee(2)
	var employee Employee
	err := db.First(&employee, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // INTE HITTAT
		return nil
	}
	return &employee
}

func Init() {
	db, _ = gorm.Open(sqlite.Open("employees.sqlite"), &gorm.Config{})
	db.AutoMigrate(&Employee{}) // Finns det en tabell i databasen som heter Employee? Om inte skapa den
	// Om det finns kolumner som inte matchar - SYNKA dom
	// Code first
	var antal int64
	db.Model(&Employee{}).Count(&antal) // Seed
	if antal == 0 {
		db.Create(&Employee{Age: 50, Namn: "Stefan", City: "Test"}) // INSERT INTO EMPLOYEES (AGE, NAMN, CITY) VALUES (50, "Stefan", "Test")
		db.Create(&Employee{Age: 14, Namn: "Oliver", City: "Test"})
		db.Create(&Employee{Age: 20, Namn: "Josefine", City: "Test"})
	}

}
