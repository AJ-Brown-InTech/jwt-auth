package main

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
)''
var secret= byte[]("auht-key")

func main(){
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

func CreateToken() (string, error){
 token := jwt.New(jwt.SigningMethodHS256)
 claim := token.Claims.(jwt.MapClaims)
 claim["exp"] = time.Now().Add(time.Hour).Unix()

 tStr, err := token.SignedString(secret)

 if err != nil{
	fmt.Println(err.Error())
	return "", err
}
return tStr, nil
}

func ValidateMe(next func(w http.ResponseWriter r* http.Request) http.Handler){
	//middleware
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header["Token"]
	})

}

func Home( w http.ResponseWriter, r *http.Request){
 fmt.Fprintf( w, "secret ")
}