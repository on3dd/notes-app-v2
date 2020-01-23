package api

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Response struct {
	Auth  bool   `json:"auth"`
	Token string `json:"token"`
	User  User   `json:"user"`
}
