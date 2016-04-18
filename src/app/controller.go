package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from %v", "Gin")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	employees := r.Group("/employees")
	employees.GET("/", func(c *gin.Context) {
		employees := Employees
		c.HTML(http.StatusOK, "employees.html",
			map[string]interface{}{
				"Employees": employees,
			})
	})
	employees.GET("/:id", func(c *gin.Context) {
		Id := c.Param("id")
		Employee, ok := Employees[Id]
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
		id := c.Param("id")
		timesOff, ok := TimesOff[id]
		if !ok {
			c.String(http.StatusNotFound, "Record Not Found")
			return
		}
		c.HTML(http.StatusOK, "vacation-overview.html",
			gin.H{
				"TimesOff": timesOff,
			})
	})

	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})
	admin.GET("/employees/:id", func(c *gin.Context) {
		if c.Param("id") == "new" {
			c.HTML(http.StatusOK, "admin-employees-new.html", nil)
		}
	})
	admin.POST("/employees/:id", func(c *gin.Context) {
		if c.Param("id") == "create" {
			Id := c.PostForm("id")
			var emp Employee
			emp.ID, _ = strconv.Atoi(c.PostForm("id"))
			emp.FirstName = c.PostForm("first_name")
			emp.LastName = c.PostForm("last_name")
			emp.Position = c.PostForm("position")
			emp.StartDate, _ = time.Parse("2006-01-02", c.PostForm("start_date"))
			Employees[Id] = emp
			c.Redirect(http.StatusMovedPermanently, "/employees/"+Id)
		}
	})

	// 1st arg: routes
	// 2nd arg: assets path
	r.Static("/public", "./public")

	return r
}
