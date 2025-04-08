package api

import (
	"log"
	"nammuru-backend/controllers"
	"nammuru-backend/jwtauth"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NammuruApi() {
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	router := gin.Default()
	config := cors.DefaultConfig()

	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Authorization", "Expire", "Token"}
	config.AllowMethods = []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}

	router.Use(cors.New(config))

	authMiddleware := jwtauth.InitJwt()
	router.POST("c1/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	p1 := router.Group("/c1")
	{
		p1.POST("/logout", authMiddleware.LogoutHandler)
		var user = new(controllers.UserRegistrationController)
		p1.POST("/register", user.Register)
	}
	c1 := router.Group("/c1")
	{
		user := new(controllers.UserRegistrationController)
		c1.GET("/otp/:phonenumber", user.OtpGeneration)
		c1.GET("/verify/:phonenumber/:otp", user.Login)
		c1.POST("/profileimage", user.AddProfileImage)

		// customer := new(controllers.CustomerController)
		// router.GET("/ws", customer.CustomerWebSocket)
	}

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
	router.Run(":8010")
}
