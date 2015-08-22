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
	s             *securecookie.SecureCookie
	cookieName    string
	cookieKeyName = "newsInstance.com"
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
	value := map[string]string{
		cookieKeyName: GenerateUniqueID(),
	}

	encoded, err := s.Encode(cookieName, value)
	if err == nil && database.SetSessionKey(encoded) {

		cookie := &http.Cookie{
			Name:  cookieName,
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		return
	}
	utils.Info("error set cookie!")
}

// ReadCookieHandler retrieve user cookie
func ReadCookieHandler(w http.ResponseWriter, r *http.Request) {
	utils.Info("reading cookie!")
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		utils.Info(fmt.Sprintf("err val of readcookie %v", err))
		utils.Info("setting new cookie")
		SetCookieHandler(w, r)
		return
	}

	utils.Info(fmt.Sprintf("cookie value, %v", cookie.Value))
	value := make(map[string]string)
	err = s.Decode(cookieName, cookie.Value, &value)
	if err != nil {
		utils.Info(fmt.Sprintf("err decoding cookie %v", err))
		return
	}
	utils.Info(fmt.Sprintf("The value of foo is %q", value[cookieKeyName]))
	return
}

// GenerateUniqueID create unique ID from bsonID
func GenerateUniqueID() string {
	objID := bson.NewObjectIdWithTime(time.Now())
	return objID.Hex()
}
