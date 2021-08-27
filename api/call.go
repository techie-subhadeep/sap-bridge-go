package api

import (
	"encoding/json"
	"sap-bridge/utils"

	"github.com/gin-gonic/gin"
	"github.com/sap/gorfc/gorfc"
)

type CallRequestBody interface{}

//
// @Summary Dispaly Connection Details
// @Description Display Connection Details
// @Produce  json
// @Router /call/{rfc} [post]
// @Param rfc path string true "RFC Name"
// @Param data body CallRequestBody true "Request Body"
// @Success 200 {string} string	"ok"
func Call(c *gin.Context) {
	sapConn, error := gorfc.ConnectionFromParams(utils.GetRfcConnectionParam())
	if error != nil {
		handleError(c, error)
		return
	}
	defer sapConn.Close()

	rfcName := c.Param("rfc")

	rawData, error := c.GetRawData()
	if error != nil {
		handleError(c, error)
		return
	}

	var jsonData map[string]interface{}
	error = json.Unmarshal(rawData, &jsonData)
	if error != nil {
		handleError(c, error)
		return
	}

	rfcResult, error := sapConn.Call(rfcName, jsonData)
	if error != nil {
		handleError(c, error)
		return
	}

	c.JSON(200, gin.H{"status": "success", "message": "", "data": rfcResult})
}
