package common

import (
	// "github.com/garyburd/redigo/redis"
	// "crypto/aes"
	// "crypto/cipher"
	// "crypto/rand"
	// "encoding/hex"
	// "os"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"log"
)

func SetUser(w http.ResponseWriter, r *http.Request, id int) {
    // fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
    // fmt.Println("YYYY-MM-DD hh:mm:ss : ", time.Now().Format("2006-01-02 15:04:05"))
    // old, _ := time.Parse("2006-01-02 15:04:05", "2016-01-02 15:04:05")
    // fmt.Println(old.Format("2006-01-02 15:04:05"))
    // now := time.Now()
    // if old.Equal(now) || old.Before(now) {  // old <= now 
    //     fmt.Println("old <= now")
    // }
	// ss := strconv.Itoa(id) + "_" + time.Now().Format("2006-01-02 15:04:05")
	txt := strconv.Itoa(id)
	fmt.Println(txt)
	txt = Encrypt("6368616e676520746869732070617373",txt)
	fmt.Println(txt)
    cookie := &http.Cookie{
        Name: "ss",
		Value: txt,
		MaxAge: 101556952,
		Secure: true,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	txt = time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(txt)
	txt = Encrypt("6368616e676520746869732070617373",txt)
	fmt.Println(txt)
    cookie1 := &http.Cookie{
        Name: "ti",
		Value: txt,
		MaxAge: 101556952,
		Secure: true,
		HttpOnly: true,
    }
    http.SetCookie(w, cookie1)
    // fmt.Fprintf(w, "Cookieの設定ができたよ")
}

func GetUser(w http.ResponseWriter, r *http.Request) int {
	usr_id := 0
    cookie, err := r.Cookie("ss")
    if err != nil {
		log.Print("No Cookie: ", err)
	}
	ss := Decrypt("6368616e676520746869732070617373",cookie.Value)
	fmt.Printf("%v\n", ss)
	delete := false
	if ss == "" {
		log.Print("ss is wrong: ", err)
		delete = true
	}
    cookie, err = r.Cookie("ti")
    if err != nil {
		log.Print("No Cookie: ", err)
	}
	t1 := Decrypt("6368616e676520746869732070617373",cookie.Value)
	fmt.Printf("%v\n", t1)
	if t1 == "" {
		log.Print("t1 is wrong: ", err)
		delete = true
	}
	if delete {
		cookie := &http.Cookie{
			Name: "ss",
			Value: "",
			MaxAge: 0,
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		cookie1 := &http.Cookie{
			Name: "ti",
			Value: "",
			MaxAge: 0,
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie1)
	}
	return usr_id
}


// func redis_set(key string, value string, c redis.Conn){
//     c.Do("SET", key, value)
//   }
  
//   func redis_get(key string, c redis.Conn) string{
//     s, err := redis.String(c.Do("GET", key))
//     if err != nil {
//       fmt.Println(err)
//       os.Exit(1)
//     }
//     return s
//   }
  
//   func redis_connection() redis.Conn {
//     const IP_PORT = "127.0.0.1:6379"
  
//     //redisに接続
//     c, err := redis.Dial("tcp", IP_PORT)
//     if err != nil {
//       panic(err)
//     }
//     return c
//   }