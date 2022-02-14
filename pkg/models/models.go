package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	UUID     string `json:"-" db:"uuid"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type SighInUser struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password_hash"`
}

type TokenClaims struct {
	UserId string
	jwt.StandardClaims
}
