package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/web"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		logger.Info("Hello, world!")
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	service := web.NewService(
		web.Name("gin-server"),
		web.Address("127.0.0.1:8888"),
		web.Handler(router),
		web.Registry(consul.NewRegistry()),
	)

	_ = service.Run()
}
