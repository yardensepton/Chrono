package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	group := router.Group("/users")
	{
		group.GET("", getUsers)
		group.GET("/:id", getUser)
	}
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{"Users": []User{{ID: "1", Name: "John", LastName: "Doe",
		Email: "john@example.com", Role: "admin"}}})
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}
	user := User{
		ID:          id,
		Name:        "Dana",
		LastName:    "Smith",
		Email:       "dana@mail.com",
		PhoneNumber: "123-456-7890",
		Role:        "user",
	}
	c.JSON(200, user)
}
