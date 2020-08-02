package common

import (
	// "fmt"
	"time"
	"net/http"
	"log"
)

func MakeCSRF(w http.ResponseWriter, r *http.Request) string {
	txt := time.Now().Format("2006-01-02 15:04:05")
	txt = Encrypt(CSRF_KEY,txt)
	cookie1 := &http.Cookie{
		Name: "xr",
		Value: txt,
		Path: "/",
		MaxAge: 120,
		Secure: true,
		HttpOnly: true,
    }
	http.SetCookie(w, cookie1)
	txt = StringRand(6)
	cookie2 := &http.Cookie{
		Name: "csrf",
		Value: txt,
		Path: "/",
		MaxAge: 120,
		Secure: true,
		HttpOnly: true,
    }
	http.SetCookie(w, cookie2)
	return txt
}

func CheckCSRF(r *http.Request, txt string) bool {
    cookie, err := r.Cookie("xr")
    value := ""
    if err == nil {
        value = cookie.Value
	}
	xr := Decrypt(CSRF_KEY,value)

	csrf, _ := time.Parse("2006-01-02 15:04:05", xr)
	t_add := csrf.Add(2 * time.Hour)

	cookie, err = r.Cookie("csrf")
    if err == nil {
        value = cookie.Value
	}

	if txt != value || time.Now().After(t_add) {
		 log.Print("CSRF Panic")
		 return false
	}
	return true
}
