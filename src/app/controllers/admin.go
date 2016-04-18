package controllers

import (
	"github.com/adlerhsieh/gin_example/src/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func registerAdmin(r *gin.Engine) {
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
			ID := createEmployee(c)
			if ID == "" {
				c.String(http.StatusBadRequest, "Failed")
				return
			}
			c.Redirect(http.StatusMovedPermanently, "/employees/"+ID)
		}
	})
}

func createEmployee(c *gin.Context) string {
	ID := c.PostForm("id")
	var emp models.Employee
	err := c.Bind(&emp)
	if err != nil {
		return ""
	}
	emp.ID, _ = strconv.Atoi(c.PostForm("id"))
	emp.StartDate, _ = time.Parse("2006-01-02", c.PostForm("start_date"))
	emp.Status = "Active"
	models.Employees[ID] = emp
	return ID
}
