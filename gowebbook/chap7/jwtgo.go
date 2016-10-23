package main

import(
	"encoding/json"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"

)

const(
	privateKeyPath = "./keys/app.rsa"
	publicKeyPath = "./keys/app.rsa.pub"
)

var(
	verifyKey *rsa.PublicKey
	signKey *rsa.PrivateKey
)

type AppClaims struct{
	UserName string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type User struct{
	UserName string `json:"username"`
	Password string `json:"password"`
}

func init(){
	var err error

	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil{
		log.Fatal("error reading private key")
		return
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil{
		log.Fatal("error reading private key")
		return
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil{
		log.Fatal("Error reading public key")
		return
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil{
		log.Fatal("error reading public key")
		return
	}
}

//reads login credentials, checks sum and creates JWT token
func loginHandler(w http.ResponseWriter, r * http.Request){
	var user User
	// decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body")
		return
	}
	//valudate user credentials
	if user.UserName != "satsi" && user.Password != "pass123"{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "sum ting wong")
		return
	}
	

	claims := AppClaims{
		user.UserName, 
		"Member",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer: "admin",
		},
	}
	log.Println(claims)
	log.Println("before token")
	//create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	
	log.Println("before key")
	log.Println(t)
	tokenString, err := t.SignedString(signKey)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while signing token")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

//only accessible with a valid token
func authHandler(w http.ResponseWriter, r *http.Request){

	//validate token
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &AppClaims{}, func(token *jwt.Token)(interface{}, error){
		//since we use only 1 private key to sign tokens and only use its public counter part to verify
		return verifyKey, nil
		})
	if err != nil{
		switch err.(type){
		case *jwt.ValidationError: //some ting wrong happened during validation
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors{
			case jwt.ValidationErrorExpired:
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, "Token expired, get new one.")
				return

			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while parsing token.")
				log.Printf("ValidationError error : %+v\n", vErr.Errors)
				return
			}
		default: //some other thing went wrong
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error parsing token!")
			log.Printf("Token parse error %v\n", err)
			return
		}
	}
	if token.Valid{
		response := Response{"Authorized to sytsem"}
		jsonResponse(response, w)
	}else{
		response := Response{"Invalid token"}
		jsonResponse(response, w)
	}
}

type Response struct{
	Text string `json:"text"`
}
type Token struct{
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter){
	json, err := json.Marshal(response)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	w.Write(json)
}

//entry point
func main(){
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth",authHandler).Methods("POST")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	log.Println("Listening....")
	server.ListenAndServe()
}