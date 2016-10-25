package common

import(
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"crypto/rsa"


)

//using asymmetric crypto/RSA keys
const(
	//openssl genrsa -out app.rsa 1024
	privateKeyPath = "keys/app.rsa"

	//openssl rsa -in app.rsa -pubout > app.rsa.pub
	publicKeyPath = "keys/app.rsa.pub"
)

//private key for signing and public key for verification
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

//Read the key files before starting handlers
func initKeys(){
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

//Generate JWT token
func GenerateJWT(name, role string)(string, error){
	claims := AppClaims{
		name, 
		role,
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
		log.Printf("Token Signing error: %v\n", err)
		return "", err
	}
	return tokenString, nil
}

//only accessible with a valid token
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc){

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
		next(w, r)
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		DisplayAppError(w, err, "Invalid Access Token", 401)
	}
}