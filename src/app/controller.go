package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	employees.GET("/:id/vacations", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]
		if !ok {
			c.String(http.StatusNotFound, "Record Not Found")
			return
		}
		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"TimesOff": timesOff,
			})
	})

	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	// 1st arg: routes
	// 2nd arg: assets path
	r.Static("/public", "./public")

	return r
}
