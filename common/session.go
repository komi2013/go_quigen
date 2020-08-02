package common

import (
	// "fmt"
	"time"
	"strconv"
	"net/http"
	"log"
)

func SetUser(w http.ResponseWriter, r *http.Request, id int) {
	txt := strconv.Itoa(id)
	txt = Encrypt(SESSION_KEY,txt)
    cookie := &http.Cookie{
        Name: "ss",
		Value: txt,
		MaxAge: 101556952,
		Secure: true,
		HttpOnly: true,
		Path: "/",
	}
	http.SetCookie(w, cookie)
	txt = time.Now().Format("2006-01-02 15:04:05")
	txt = Encrypt(SESSION_KEY,txt)
    cookie1 := &http.Cookie{
        Name: "ti",
		Value: txt,
		MaxAge: 101556952,
		Secure: true,
		HttpOnly: true,
		Path: "/",
    }
    http.SetCookie(w, cookie1)
}

func GetUser(w http.ResponseWriter, r *http.Request) int {
	usr_id := 0
    cookie, err := r.Cookie("ss")
    if err != nil {
		log.Print("No Cookie: ", err)
		return 0
	}
	ss := Decrypt(SESSION_KEY,cookie.Value)

	delete := false
	if ss == "" {
		log.Print("ss is wrong: ", err)
		delete = true
	}
	usr_id, err = strconv.Atoi(ss)
	if err != nil {
		log.Print("ss error: ", err)
	}
    cookie, err = r.Cookie("ti")
    if err != nil {
		log.Print("No Cookie: ", err)
	}
	t1 := Decrypt(SESSION_KEY,cookie.Value)

	if t1 == "" {
		log.Print("t1 is wrong: ", err)
		delete = true
	}
	stampTime, err := time.Parse("2006-01-02 15:04:05", t1)
	
	if err != nil {
		log.Print("time.Parse error: ", err)
	}
	t1_add := stampTime.AddDate(1,0,0)
	if time.Now().After(t1_add) {
		log.Print("session expired")
		delete = true
	}
	if delete {
		cookie := &http.Cookie{
			Name: "ss",
			Value: "",
			MaxAge: 0,
			Secure: true,
			HttpOnly: true,
			Path: "/",
		}
		http.SetCookie(w, cookie)
		cookie1 := &http.Cookie{
			Name: "ti",
			Value: "",
			MaxAge: 0,
			Secure: true,
			HttpOnly: true,
			Path: "/",
		}
		http.SetCookie(w, cookie1)
		return 0
	}
	return usr_id
}