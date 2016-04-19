package controllers

import (
	"fmt"
	"github.com/adlerhsieh/gin_example/src/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerEmployees(r *gin.Engine) {
	employees := r.Group("/employees")
	employees.GET("/", func(c *gin.Context) {
		employees := models.Employees
		c.HTML(http.StatusOK, "employees.html",
			map[string]interface{}{
				"Employees": employees,
			})
	})
	employees.GET("/:id", func(c *gin.Context) {
		ID := c.Param("id")
		Employee, ok := models.Employees[ID]
		if !ok {
			c.String(http.StatusNotFound, "Record Not Found")
			return
		}
		c.HTML(http.StatusOK, "employee.html",
			gin.H{
				"Employee": Employee,
			})
	})
	employees.GET("/:id/vacations", func(c *gin.Context) {
		// "/add" will be considered an {id: "add"}
		ID := c.Param("id")
		timesOff, ok := models.TimesOff[ID]
		if !ok {
			c.String(http.StatusNotFound, "Record Not Found")
			return
		}
		c.HTML(http.StatusOK, "vacation-overview.html",
			gin.H{
				"ID":       ID,
				"TimesOff": timesOff,
			})
	})
	employees.POST("/:id/vacations", func(c *gin.Context) {
		var timeOff models.TimeOff
		err := c.BindJSON(&timeOff)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "Failed")
			return
		}
		ID := c.Param("id")
		timesOff, ok := models.TimesOff[ID]
		if !ok {
			models.TimesOff[ID] = []models.TimeOff{}
		}
		models.TimesOff[ID] = append(timesOff, timeOff)
		c.String(http.StatusOK, "success")
	})
}
