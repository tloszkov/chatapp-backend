package main

import (
	"ChatApp/Src/Routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()

	Routes.UserRoutes(router)
	Routes.MessageRoutes(router)
	Routes.GroupMessageRoutes(router)
	Routes.MessageBoardRoutes(router)

	router.Run(":8090")
}
