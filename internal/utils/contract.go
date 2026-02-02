package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateContractID(category string) string {
	var cat string
	switch category {
	case "MOBIL":
		cat = "MBL"
	case "MOTOR":
		cat = "MTR"
	case "WHITE_GOODS":
		cat = "WG"
	default:
		cat = "INV"
	}

	date := time.Now().Format("20060102")

	uniquePart := uuid.New().String()[:8]

	nomor_kontrak := fmt.Sprintf("%s-%s-%s", cat, date, uniquePart)

	return nomor_kontrak
}
