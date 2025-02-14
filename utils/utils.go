package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Utilization struct {
	BaseUrlGraphApi  string
	UserID           string
	SecretAPP        string
	AppID            string
	SystemUserAccess string
}

func GetEnv() (result Utilization) {
	workDir, _ := os.Getwd()
	loadEnv := godotenv.Load(fmt.Sprintf("%s/.env", workDir))
	if loadEnv == nil {
		// GRAPH API
		BaseUrlGraphApi := os.Getenv("BASE_URL_GRAPH_API")
		UserId := os.Getenv("USER_ID")
		SecretApp := os.Getenv("SECRET_APP")
		AppId := os.Getenv("APP_ID")
		SystemUserAccess := os.Getenv("SYSTEM_USER_ACCESS")

		return Utilization{
			BaseUrlGraphApi:  BaseUrlGraphApi,
			UserID:           UserId,
			SecretAPP:        SecretApp,
			AppID:            AppId,
			SystemUserAccess: SystemUserAccess,
		}
	}
	return
}

var Utils = GetEnv()
