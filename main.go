package main

import (
	"ChatApp/Src/DBConnector"
	"ChatApp/Src/Middleware"
	"ChatApp/Src/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CorsPolicy)

	logger, logFile := middleware.Logg("app")
	defer logFile.Close()

	router.Use(middleware.LoggerMiddleware(logger))

	DBConnector.ConnectToMongo()
	defer DBConnector.DisconnectFromMongo()
	client := DBConnector.GetMongoClient()

	Routes.UserRoutes(router, client)
	Routes.MessageRoutes(router, client)
	Routes.GroupMessageRoutes(router, client)
	Routes.MessageBoardRoutes(router, client)

	router.Run(":8090")
}
