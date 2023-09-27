package cookie

import (
	"net/http"
	"time"
)

const SID = "token"

func SetSID(val string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = SID
	cookie.Value = val
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	return cookie
}

func DelSID() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = SID
	cookie.Path = "/"
	cookie.MaxAge = -1
	return cookie
}
