package common

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//using asymmetric crypto/RSA keys
const (
	//openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"

	//openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

//Private key for signing and public key for verification
var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}	
}

func GenerateJWT(anme, role string) (string,error) {
	//Create a signer for rsa 256
	t := jwt.New(jwt.GetSigningMethod('RS256'))

	//Set claim for jwt token
	t.Claims["iss"] = "admin"
	t.Claims["UserInfo"] = struct {
		Name string
		Role string
	}{name, role}

	//Set the expiry tiem for jwt token
	t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "",err
	} 
	return tokenString, nil
}

//Middleware for validating jwt tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//Validate the token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		//Verify token with public key which is the couter part of private key
		return verifyKey, nil
	}

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired: 
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new token",
					401
				)
				return

			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token!",
					500
				)
				return
				
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing the Access Token!",
				500
			)
			return
			
		}
	}
	if token.Valid {
		next(w,r)
	} else {
		DisplayAppError(
			w,
			err,
			"Invalid Access Token!",
			401
		)
	}
}