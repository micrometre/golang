// routes/users.go

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", getUsers)
	}
}
func getUsers(c *gin.Context) {
	// ... your logic to get all users
	c.JSON(http.StatusOK, gin.H{"data": "Get all users"})
}
