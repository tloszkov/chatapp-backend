package main

import (
	"ChatApp/Src/DBConnector"
	"ChatApp/Src/Middleware"
	"ChatApp/Src/Routes"
	_ "ChatApp/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title ChatApp API
// @version 1.0
// @description This is a sample server for a chat application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8090
// @BasePath /api/v1/

func main() {
	router := gin.Default()
	router.Use(middleware.CorsPolicy)

	logger, logFile := middleware.Logg("app")
	defer logFile.Close()

	router.Use(middleware.LoggerMiddleware(logger))
	//router.Use(middleware.AuthMiddleware())
	router.Use(middleware.RateLimitMiddleware())

	DBConnector.ConnectToMongo()
	defer DBConnector.DisconnectFromMongo()
	client := DBConnector.GetMongoClient()

	Routes.UserRoutes(router, client)
	Routes.MessageRoutes(router, client)
	Routes.GroupMessageRoutes(router, client)
	Routes.MessageBoardRoutes(router, client)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8090")
}
