package cookie

import (
	"clean-architecture/config"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

const SID = "token"

type CookieManager struct {
	sc *securecookie.SecureCookie
}

func NewCookieManager(conf *config.Config) *CookieManager {
	return &CookieManager{sc: securecookie.New([]byte(conf.HashKey), []byte(conf.BlockKey))}
}

func (cm *CookieManager) SetSID(val string) (*http.Cookie, error) {
	encoded, err := cm.sc.Encode(SID, val)
	if err != nil {
		return nil, err
	}
	cookie := new(http.Cookie)
	cookie.Name = SID
	cookie.Value = encoded
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	return cookie, nil
}

func (cm *CookieManager) GetUserSID(getCookie *http.Cookie) (string, error) {
	var userSID string
	if err := cm.sc.Decode(SID, getCookie.Value, &userSID); err != nil {
		return "", err
	}
	return userSID, nil
}

func DelSID() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = SID
	cookie.Path = "/"
	cookie.MaxAge = -1
	return cookie
}
