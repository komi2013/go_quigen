package common

import (
	// "fmt"
	"time"
	"net/http"
	"log"
	"github.com/catinello/base62"
	// "strconv"
)

func MakeCSRF(w http.ResponseWriter, r *http.Request) string {
	u62 := base62.Encode( int(time.Now().Unix()) )
	page := Encrypt(CSRF_KEY,u62)
	cookie := &http.Cookie{
		Name: "xr",
		Value: u62,
		Path: "/",
		MaxAge: 0,
		Secure: true,
		HttpOnly: true,
    }
	http.SetCookie(w, cookie)
	return page
}

func CheckCSRF(r *http.Request, page string) bool {
    cookie, err := r.Cookie("xr")
    u62 := ""
    if err == nil {
        u62 = cookie.Value
	}
	if Decrypt(CSRF_KEY,page) != u62 || u62 == "" {
		log.Print("CSRF page cookie is not match")
		return false
	}
	n, err := base62.Decode(u62)
    if err != nil {
		log.Print(err)
	}
    tm := time.Unix(int64(n), 0)
	t_add := tm.Add(2 * time.Hour)
	if time.Now().After(t_add) {
		log.Print("CSRF is expired")
		return false
	}
	return true
}
