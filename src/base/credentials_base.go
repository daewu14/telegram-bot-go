package base

import (
	"errors"
	"log"
)

// baseUrl is variable string to cache base url of telegram
var baseUrl = "https://api.telegram.org/"

// _token is variable string to cache _token string of telegram
var _token = ""

type Error struct {
}

// GetToken is Getter to getting _token cached
func GetToken() string {
	if _token == "" {
		log.Panic("Token not found")
		panic(errors.New("Token not found"))
	}
	return _token
}

// SetToken is Setter to set the token
func SetToken(token string) { _token = token }

// GetBaseUrl is Getter to getting base url
func GetBaseUrl() string {
	if _token != "" {
		baseUrl += "bot"+_token+"/"
	}
	return baseUrl
}