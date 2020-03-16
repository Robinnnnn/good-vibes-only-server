package utils

import (
	"net/http"
	"time"
)

// AddCookie applies a cookie to the response of an HTTP request
func AddCookie(w http.ResponseWriter, key string, value string) {
	expiry := time.Now().Add(time.Minute * 10)
	cookie := http.Cookie{
		Name:    key,
		Value:   value,
		Expires: expiry,
	}
	http.SetCookie(w, &cookie)
}
