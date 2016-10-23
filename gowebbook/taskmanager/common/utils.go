package common

import(
	"encoding/json"
	"log"
	"os"
)

type configuration struct{
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

//AppConfig to store app related configs

var AppConfig configuration

//Initialize app config
func initConfig(){
	loadAppConfig()
}

//Read config.json & decode AppConfig
func loadAppConfig(){
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig] :%s\n", err)
	}
}