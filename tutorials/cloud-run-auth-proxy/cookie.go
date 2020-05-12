package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

var cookieSecret []byte
var cookieName = "AUTH_SESSION"
var s *securecookie.SecureCookie

func init() {
	value, exists := os.LookupEnv("COOKIE_SECRET")
	if !exists {
		panic("You need to set a COOKIE_SECRET env var to sign auth session cookie")
	}
	cookieSecret = []byte(value)
	s = securecookie.New(cookieSecret, nil)
}

func setCookie(w http.ResponseWriter, userEmail string) {
	encoded, err := s.Encode(cookieName, userEmail)
	if err == nil {
		cookie := &http.Cookie{
			Name:     cookieName,
			Value:    encoded,
			Path:     "/",
			MaxAge:   int(time.Duration(2 * time.Hour).Seconds()),
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)
	}
}

func readCookie(r *http.Request) (userEmail string, err error) {
	var cookie *http.Cookie
	if cookie, err = r.Cookie(cookieName); err == nil {
		var value string
		if err = s.Decode(cookieName, cookie.Value, &value); err == nil {
			fmt.Println("decoded", err)
			return value, err
		}
		return "", err

	}

	return "", err

}
