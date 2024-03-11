package main

import (
	"github.com/gin-gonic/gin"
	"go_test/webook/internal/web"
)

func main() {
	server := gin.Default()
	//v1 := server.Group("/v1")
	//users := server.Group("/users/v1")
	u := web.NewUserHandler()
	//u.RegisterRoutesV1(server.Group("/users"))
	u.RegisterRoutes(server)
	server.Run(":8080")
}
