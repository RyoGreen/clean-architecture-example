package cookie

import (
	"net/http"
	"time"
)

const name = "token"

func SetTokenCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	return cookie
}

func DeleteTokenCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Path = "/"
	cookie.MaxAge = -1
	return cookie
}
