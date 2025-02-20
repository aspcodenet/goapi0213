package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"systementor.se/goapi0213/data"
)

func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id") // "2"
	numId, _ := strconv.Atoi(id)
	employee := data.GetEmployee(numId)

	if employee == nil { // INTE HITTAT  /api/employee/
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func apiEmployeeAdd(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id = 0
	data.CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)

}

func apiEmployeeUpdateById(c *gin.Context) {
	id := c.Param("id")
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	employee.Id, _ = strconv.Atoi(id)
	if data.UpdateEmployee(employee) == false {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

type PageView struct {
	Title  string
	Rubrik string
}

func handleStartPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", &PageView{Title: "Hello World", Rubrik: data.GetAllEmployees()[0].Namn})
	//c.String(http.StatusOK, "Hello World")
}

var config Config

func main() {
	// PÅ TORSDAG
	// Jag vill köra Mysql
	// eller sqlite men filemn - eomplyees.sqhej

	// CONFIG -> flera lager med konfigurering
	//    	config filer (vi kommer att skriva dom i YAML - application.properties,
	//								settings-fil (JSON)
	// 		environment-variabler (VIKTIGT FÖR CONTAINERS)
	readConfig(&config)

	fmt.Println("Värdet i database file")
	fmt.Println(config.Database.File)

	data.Init(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	router := gin.Default()
	router.LoadHTMLGlob("templates/**")

	// HTTP - URL ? Men också vilket VERB

	// router - skicka vidare beroende på   på address
	router.GET("/", handleStartPage)                      // READ
	router.GET("/api/employee", handleGetAllEmployees)    // READ - sortera, filtrera, paginera
	router.GET("/api/employee/:id", handleGetOneEmployee) // READ

	// Delete gör vi sällan!!!
	// OFTARE SOFT DELETE  - update product set active=false where id=12
	// db.Delete(employee)

	// UPDATE, DELETE, CREATE - dom går inte att testa/anropa via en webbläsare
	router.POST("/api/employee", apiEmployeeAdd) // POSTA JSON för en Employee till /api/employee
	router.PUT("/api/employee/:id", apiEmployeeUpdateById)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
