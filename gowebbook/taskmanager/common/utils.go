package common

import(
	"log"
	"os"
	"net/http"
	"encoding/json"
)

type configuration struct{
	Server, MongoDBHost, DBUser, DBPwd, Database string
}

type (
        appError struct {
                Error      string `json:"error"`
                Message    string `json:"message"`
                HttpStatus int    `json:"status"`
        }
        errorResource struct {
                Data appError `json:"data"`
} )

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
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	log.Println(AppConfig)
	if err != nil{
		log.Fatalf("[loadAppConfig]:%s\n",err)
	}
}

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int){
	errObj := appError{
		Error: handlerError.Error(),
		Message: message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil{
		w.Write(j)
	}
}