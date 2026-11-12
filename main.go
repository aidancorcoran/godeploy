package main

import (
	"github.com/aidancorcoran/godeploy/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Api V1
	v1 := router.Group("/api/v1")
	{
		v1.GET("health", api.GetHealth)
	}

	router.Run()
}
