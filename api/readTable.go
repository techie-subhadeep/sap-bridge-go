package api

import (
	"encoding/json"
	"sap-bridge/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/sap/gorfc/gorfc"
)

type ReadTableRequestBody struct {
	Fields   []string `json:"fields"`
	Where    []string `json:"where"`
	MaxRows  int      `json:"maxRows"`
	SkipRows int      `json:"skipRows"`
}

type rfcReadTableOptionBody struct {
	TEXT string
}

type rfcReadTableFields struct {
	FIELDNAME string
	OFFSET    string
	LENGTH    string
	T         string
	FIELDTEXT string
}
type rfcReadTableFuncBody struct {
	QUERY_TABLE string
	DELIMITER   string
	NO_DATA     string
	ROWSKIPS    int
	ROWCOUNT    int
	OPTIONS     []rfcReadTableOptionBody
	FIELDS      []rfcReadTableFields
}

type rfcReadTableOtput struct {
	FIELDS []rfcReadTableFields
	DATA   []struct{ WA string }
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
	sapConn, error := gorfc.ConnectionFromParams(utils.GetRfcConnectionParam())
	if error != nil {
		handleError(c, error)
		return
	}
	defer sapConn.Close()

	const delimiter = "|"
	tableName := c.Param("table")
	rawData, _ := c.GetRawData()

	jsonBody := new(ReadTableRequestBody)
	json.Unmarshal(rawData, &jsonBody)

	options := make([]rfcReadTableOptionBody, len(jsonBody.Where))
	for i := 0; i < len(jsonBody.Where); i++ {
		options[i].TEXT = jsonBody.Where[i]
	}

	fields := make([]rfcReadTableFields, len(jsonBody.Fields))
	for i := 0; i < len(jsonBody.Fields); i++ {
		fields[i].FIELDNAME = jsonBody.Fields[i]
	}

	rfcBody := rfcReadTableFuncBody{
		QUERY_TABLE: tableName,
		DELIMITER:   delimiter,
		NO_DATA:     "",
		ROWSKIPS:    jsonBody.MaxRows,
		ROWCOUNT:    jsonBody.SkipRows,
		OPTIONS:     options,
		FIELDS:      fields,
	}

	tableRawData, error := sapConn.Call("RFC_READ_TABLE", rfcBody)
	if error != nil {
		handleError(c, error)
		return
	}
	tableData := new(rfcReadTableOtput)
	mapstructure.Decode(tableRawData, &tableData)

	opFieldName := make([]string, len(tableData.FIELDS))
	for i := 0; i < len(tableData.FIELDS); i++ {
		opFieldName[i] = tableData.FIELDS[i].FIELDNAME
	}

	result := make([]map[string]string, len(tableData.DATA))

	for i := 0; i < len(tableData.DATA); i++ {
		dataSplit := strings.Split(tableData.DATA[i].WA, delimiter)
		r := make(map[string]string)
		for idx, val := range dataSplit {
			r[opFieldName[idx]] = strings.TrimSpace(val)
		}
		result[i] = r
	}
	c.JSON(200, gin.H{"status": "success", "data": result})
}
