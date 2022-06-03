package main

import (
	"fmt"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/apis"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", apis.GetUser)
	}
	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
	// load application configurations

}
