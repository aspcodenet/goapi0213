package main

import (
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

func main() {
	data.Init()

	router := gin.Default()

	// router - skicka vidare beroende på   på address
	router.GET("/api/employee", handleGetAllEmployees)    // READ
	router.GET("/api/employee/:id", handleGetOneEmployee) // READ
	// UPDATE, DELETE, CREATE - dom går inte att testa/anropa via en webbläsare

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
