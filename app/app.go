package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tom-camp/api.tmc/app/controllers"
	"github.com/tom-camp/api.tmc/app/middlewares"
	"go.uber.org/zap"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		zap.S().Info("No .env file found")
	}
}

func main() {
	// Init gin router
	router := gin.Default()

	// Its great to version your API's
	v1 := router.Group("/api/v1")
	{
		// Define the hello controller
		hello := new(controllers.HelloController)
		// Define a GET request to call the Default
		// method in controllers/hello.go
		v1.GET("/hello", hello.Default)

		// Define the user controller
		user := new(controllers.UserController)

		// Define the link controller
		link := new(controllers.BookmarkController)

		// Create the signup endpoint
		v1.POST("/signup", user.Signup)
		// Create the login endpoint
		v1.POST("/login", user.Login)
		// Password reset
		v1.PUT("/password-reset", user.PasswordReset)
		// Send reset link
		v1.PUT("/reset-link", user.ResetLink)
		// Send verify link
		v1.PUT("/verify-link", user.VerifyLink)
		// Verify account
		v1.PUT("/verify-account", user.VerifyAccount)
		// Refresh token
		v1.GET("/refresh", user.RefreshToken)

		bookmarks := v1.Group("/bookmarks")

		bookmarks.Use(middlewares.Authenticate())

		{
			bookmarks.GET("/all", link.FetchBookmarks)
			bookmarks.POST("/create", link.CreateBookmak)
			bookmarks.DELETE("/delete", link.DeleteBookmark)
		}
	}

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// Init our server
	router.Run(":5000")
}
