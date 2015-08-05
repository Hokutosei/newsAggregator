package security

import (
	"fmt"
	"net/http"
	"threadtimer/lib/utils"

	"github.com/gorilla/securecookie"
)

var (
	hashKey       = []byte{}
	blockKey      = []byte{}
	s             *securecookie.SecureCookie
	cookieName    string
	cookieKeyName = "username"
)

// BuildSecureKeys start building required keys
func BuildSecureKeys(hash, block, cookie string) {
	hashKey = []byte(hash)
	blockKey = securecookie.GenerateRandomKey(32)
	s = securecookie.New(hashKey, blockKey)
	cookieName = cookie

	utils.Info("built secure cookies!")
}

// SetCookieHandler set cookie handler to user
func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info("setcookieHandler!")
	value := map[string]string{
		cookieKeyName: "jeanepaul",
		"password":    "fdsafsadfadsfsadfasfs",
	}
	encoded, err := s.Encode(cookieName, value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

// ReadCookieHandler retrieve user cookie
func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie(cookieName); err == nil {
		value := make(map[string]string)
		if err = s.Decode(cookieName, cookie.Value, &value); err == nil {
			// fmt.Fprintf(w, "The value of foo is %q", value[cookieKeyName])
			utils.Info(fmt.Sprintf("The value of foo is %q", value[cookieKeyName]))
		}
	}
}
