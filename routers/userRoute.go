package routeres

import (
	"github.com/gin-gonic/gin"
	"my-go-server/users"
	"my-go-server/services"
)

func RegisterUserRoutes(router *gin.Engine, userService *services.UserService) {
	group := router.Group("/users")
	{
		group.GET("", func(c *gin.Context) {getUsers(c, userService)})
		group.GET("/:id", func(c *gin.Context) {getUser(c, userService)})
		group.POST("/", func(c *gin.Context) {createNewUser(c, userService)})
	}
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context, userService *services.UserService) {
	
	c.JSON(200, gin.H{"Users": []users.User{{ID: "1", Name: "John", LastName: "Doe",
		Email: "john@example.com", Role: "admin"}}})
}

// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} User
// @Router /users/{id} [get]
func getUser(c *gin.Context, userService *services.UserService) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "ID is required"})
		return
	}
	user, err := userService.GetUserByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}
	// user := users.User{
	// 	ID:          id,
	// 	Name:        "Dana",
	// 	LastName:    "Smith",
	// 	Email:       "dana@mail.com",
	// 	PhoneNumber: "123-456-7890",
	// 	Role:        "user",
	// }
	c.JSON(200, user)
}

// @Summary Create a user
// @Description Create a new user
// @Tags users
// @Param user body UserRequest true "User details"
// @Produce json
// @Success 201 {object} User
// @Router /users [post]
func createNewUser(c *gin.Context, userService *services.UserService) {
	var newUser users.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdUser, err := userService.InsertUser(newUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, createdUser)
}