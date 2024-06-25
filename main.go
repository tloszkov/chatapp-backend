package main

import (
	"ChatApp/Src/DBConnector"
	"ChatApp/Src/Routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()

	DBConnector.ConnectToMongo()
	defer DBConnector.DisconnectFromMongo()
	client := DBConnector.GetMongoClient()

	Routes.UserRoutes(router, client)
	Routes.MessageRoutes(router, client)
	Routes.GroupMessageRoutes(router)
	Routes.MessageBoardRoutes(router, client)

	router.Run(":8090")

}
