package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
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

	registerEmployees(r)
	registerAdmin(r)

	// 1st arg: routes
	// 2nd arg: assets path
	r.Static("/public", "./public")
}
