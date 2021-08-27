package main

import (
	"sap-bridge/api"
	_ "sap-bridge/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title SAP RFC Bridge
// @version 1.0
// @description This project provide REST API to intract with SAP RFC
// @termsOfService http://swagger.io/terms/

// @contact.name Developer
// @contact.url https://github.com/subhadeepdas91
func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	r.GET("/ping", api.Ping)
	r.POST("/call/:rfc", api.Call)
	r.POST("/read-table/:table", api.ReadTable)
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:80
}
