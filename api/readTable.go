package api

import (
	"sap-bridge/utils"

	"github.com/gin-gonic/gin"
	"github.com/sap/gorfc/gorfc"
)

type ReadTableRequestBody struct {
	Fields   []string `json:"fields"`
	Where    []string `json:"where"`
	MaxRows  int      `json:"maxRows"`
	SkipRows int      `json:"skipRows"`
}

//
// @Summary Read Table
// @Description Read ERP Table Data
// @Produce  json
// @Router /read-table/{table} [post]
// @Param table path string true "Table Name"
// @Param data body ReadTableRequestBody true "Request Body"
// @Success 200 {string} string	"ok"
func ReadTable(c *gin.Context) {
	sapConn, _ := gorfc.ConnectionFromParams(utils.GetRfcConnectionParam())
	if sapConn != nil {
		defer sapConn.Close()
	}

	tableName := c.Param("table")
	// if error != nil {
	// 	c.JSON(400, gin.H{"status": "error", "message": error.Error()})
	// }
	c.JSON(200, gin.H{"status": "success", "data": tableName})
}
