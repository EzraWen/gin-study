package main

import "github.com/gin-gonic/gin"
import . "com.ezra/gin-test/api"

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexApi)
	router.POST("/person/add", AddPersonApi)
	return router
}
