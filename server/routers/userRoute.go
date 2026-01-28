package routers

import (
	"my-go-server/middleware"
	"my-go-server/model/users"
	"my-go-server/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userService *services.UserService,authService *services.AuthService) {
	group := router.Group("/users")
	{
		// Public routes (no authentication required)
		group.POST("/login", func(c *gin.Context) { loginUser(c, authService) })
		group.POST("/", func(c *gin.Context) { createNewUser(c, userService) })
		group.POST("/register", func(c *gin.Context) { registerUser(c, userService) })
		group.POST("/refresh-token", func(c *gin.Context) { refreshToken(c, authService, userService) })

		// Protected routes (authentication required)
		protected := group.Group("")
		protected.Use(middleware.JWTMiddleware())
		{
			protected.GET("", func(c *gin.Context) { getUsers(c, userService) })
			protected.GET("/:id", func(c *gin.Context) { getUser(c, userService) })
		}
	}
}

// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} users.User
// @Security BearerAuth
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
// @Success 200 {object} users.User
// @Security BearerAuth
// @Router /users/{id} [get]
func getUser(c *gin.Context, userService *services.UserService) {
	id := c.Param("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "invalid_request"})
		return
	}
	user, err := userService.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user_not_found"})
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
// @Param user body users.UserRequest true "User details"
// @Produce json
// @Success 201 {object} users.User
// @Router /users [post]
func createNewUser(c *gin.Context, userService *services.UserService) {
	var newUser users.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}
	createdUser, err := userService.InsertUser(newUser)
	if err != nil {
		c.JSON(500, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(201, gin.H{"message": "user_created", "user": createdUser})
}

// @Summary Refresh access token
// @Description Refresh JWT access token using a refresh token
// @Tags auth
// @Param refresh_token body users.RefreshRequest true "Refresh token"
// @Produce json
// @Success 200 {object} map[string]string
// @Router /users/refresh-token [post]
func refreshToken(c *gin.Context, authService *services.AuthService, userService *services.UserService) {
	var refreshRequest users.RefreshRequest
	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(400, gin.H{"error": "invalid_request"})
		return
	}

	newAccessToken, newRefreshToken, err := authService.RefreshTokens(refreshRequest.RefreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"access_token": newAccessToken, "refresh_token": newRefreshToken})

}

// @Summary Register a new user
// @Description Register a new user account
// @Tags auth
// @Param user body users.UserRequest true "User registration details"
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /users/register [post]
func registerUser(c *gin.Context, userService *services.UserService) {
	var newUser users.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}

	createdUser, err := userService.InsertUser(newUser)
	if err != nil {
		c.JSON(500, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(201, gin.H{"message": "user_registered", "user": createdUser})
}

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Param credentials body users.LoginRequest true "User login credentials"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/login [post]
func loginUser(c *gin.Context, AuthService *services.AuthService) {
	var loginReq users.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": gin.H{"error": err.Error()}})
		return
	}

	user, accessToken, refreshToken, err := AuthService.AuthenticateUser(loginReq.Email, loginReq.Password)

	if err != nil {
		c.JSON(401, gin.H{"error": "invalid_credentials"})
		return
	}
	
	c.JSON(200, gin.H{
		"message": "login_successful",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
