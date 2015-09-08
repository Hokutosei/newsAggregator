package security

import (
	"fmt"
	"net/http"
	"threadtimer/lib/utils"
	"time"
	"web_apps/news_aggregator/modules/database"
	"web_apps/news_aggregator/modules/httpHandlers"

	"github.com/gorilla/securecookie"
	"labix.org/v2/mgo/bson"
)

var (
	hashKey       = []byte{}
	blockKey      = []byte{}
	rBkey         = []string{"secured", "blockKey"}
	s2            *securecookie.SecureCookie
	cookieName    string
	cookieKeyName = "newsInstance.com"
)

// BuildSecureKeys start building required keys
func BuildSecureKeys(hash, block, cookie string) {
	getBlockKey := make(chan []byte)
	go blockKeyGetter(getBlockKey)
	hashKey = []byte(hash)
	blockKey = <-getBlockKey
	s2 = securecookie.New(hashKey, blockKey)
	cookieName = cookie

	utils.Info("built secure cookies!")
}

// BuildHTTPObjectIDKey generate unique bsonObject ID
func BuildHTTPObjectIDKey(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"sessionId": GenerateUniqueStrID()}
	httpHandlers.PublicrespondToJSON(w, response)
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

// blockKeyGetter retrieve block key from redis if exists
// else generate randomkey and register to redis
func blockKeyGetter(bKeyChan chan []byte) {
	bKey := database.Rstring{database.RedisKeyGen(rBkey...), ""}
	secureCookieKey, err := bKey.Get()
	if err != nil {
		utils.Info(fmt.Sprintf("err blockkeygetter %v", err))
		bKeyChan <- registerSecureCookieKey()
	}
	bKeyChan <- []byte(secureCookieKey.(string))
}

// registerSecureCookieKey register new random key to redis
func registerSecureCookieKey() []byte {
	val := securecookie.GenerateRandomKey(32)
	key := database.Rstring{database.RedisKeyGen(rBkey...), string(val)}
	_, err := key.Set()
	if err != nil {
		panic(err)
	}
	return val
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

// GenerateUniqueStrID unique string ID
func GenerateUniqueStrID() string {
	objID := bson.NewObjectIdWithTime(time.Now())
	return objID.String()
}
