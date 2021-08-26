package api

import (
	"sap-bridge/utils"

	"github.com/gin-gonic/gin"
	"github.com/sap/gorfc/gorfc"
)

//
// @Summary Dispaly Connection Details
// @Description Display Connection Details
// @Produce  json
// @Router /ping [get]
// @Success 200 {string} string	"ok"
func Ping(c *gin.Context) {
	sapConn, _ := gorfc.ConnectionFromParams(utils.GetRfcConnectionParam())
	if sapConn != nil {
		defer sapConn.Close()
	}
	connAttr, _ := sapConn.GetConnectionAttributes()
	c.JSON(200, gin.H{
		"message": connAttr,
	})
}
