package utils

import (
	"os"

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
