package main

import (
	"github.com/gin-gonic/examples/group-routes/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/", "./public")
	routes.Run()
}
