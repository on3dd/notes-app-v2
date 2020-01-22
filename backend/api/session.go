package api

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func setSession(userName string, w http.ResponseWriter) error {
	value := map[string]string{"name": userName}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Expires: time.Now().Add(24 * time.Hour),
		}
		http.SetCookie(w, cookie)
	} else {return err}

	return nil
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
