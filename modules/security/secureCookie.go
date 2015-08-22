package security

import (
	"fmt"
	"net/http"
	"threadtimer/lib/utils"
	"time"
	"web_apps/news_aggregator/modules/database"

	"github.com/gorilla/securecookie"
	"labix.org/v2/mgo/bson"
)

var (
	hashKey       = []byte{}
	blockKey      = []byte{}
	s2            *securecookie.SecureCookie
	cookieName    string
	cookieKeyName = "newsInstance.com"
)

// BuildSecureKeys start building required keys
func BuildSecureKeys(hash, block, cookie string) {
	hashKey = []byte(hash)
	blockKey = securecookie.GenerateRandomKey(32)
	s2 = securecookie.New(hashKey, blockKey)
	cookieName = cookie

	utils.Info("built secure cookies!")
}

// SetCookieHandler set cookie handler to user
func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value := map[string]string{
		cookieKeyName: GenerateUniqueID(),
	}

	encoded, err := s2.Encode(cookieName, value)
	if err == nil && database.SetSessionKey(encoded) {

		cookie := &http.Cookie{
			Name:   cookieName,
			Value:  encoded,
			Path:   "/",
			MaxAge: 0,

			// addDate(year, month, day)
			Expires: time.Now().AddDate(0, 1, 0),
		}
		http.SetCookie(w, cookie)
		return
	}
	utils.Info("error set cookie!")
}

// ReadCookieHandler retrieve user cookie
func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		SetCookieHandler(w, r)
		return
	}

	// utils.Info(fmt.Sprintf("cookie value, %v", cookie.Value))
	value := make(map[string]string)
	utils.Info(fmt.Sprintf("cookie read %v", cookie.Value))
	err = s2.Decode(cookieName, cookie.Value, &value)
	if err != nil {
		utils.Info(fmt.Sprintf("err decoding cookie %v", err))
		return
	}
	utils.Info(fmt.Sprintf("The value of foo is %q", value[cookieKeyName]))
	registerSessionID(value[cookieKeyName])
	return
}

func registerSessionID(cookieVal string) {
	if cookieVal != "" && GetSessionID(cookieVal) {
		utils.Info("session found!")
		return
	}
	utils.Info("session not found!")
}

// GenerateUniqueID create unique ID from bsonID
func GenerateUniqueID() string {
	objID := bson.NewObjectIdWithTime(time.Now())
	return objID.Hex()
}
