package main

import (
	"github.com/adlerhsieh/gin_example/src/app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.RegisterRoutes(r)
	r.Run(":3000")
}
