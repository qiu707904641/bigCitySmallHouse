package main

import (
	"bigCitySmallHouse/mongodb"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	opts := mongodb.NewOptions()
	opts.Uri = "mongodb://admin:admin@43.138.174.42:27017/"
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/publish", Publish)
		apiGroup.GET("/publish/list", PublishList)
	}
	err := engine.Run(":10002")
	if err != nil {
		log.Fatalln(err)
	}
}
