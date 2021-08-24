package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sap/gorfc/gorfc"
)

func GetRfcConnectionParam() gorfc.ConnectionParameters {
	return gorfc.ConnectionParameters{
		"Client": os.Getenv("SAP_CLIENT"),
		"User":   os.Getenv("SAP_USER"),
		"Passwd": os.Getenv("SAP_PASSWD"),
		"Ashost": os.Getenv("SAP_ASHOST"),
		"Sysnr":  os.Getenv("SAP_SYSNR"),
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		sapConn, _ := gorfc.ConnectionFromParams(GetRfcConnectionParam())
		if sapConn != nil {
			defer sapConn.Close()
		}
		connAttr, _ := sapConn.GetConnectionAttributes()
		c.JSON(200, gin.H{
			"message": connAttr,
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:80
}
