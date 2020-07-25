package app

import (
    "net/http"
    "github.com/garyburd/redigo/redis"
    "fmt"
    "os"
    "log"
    // "time"
    "../common"
    // "net/url"
    "strings"
)

func CryptTry(w http.ResponseWriter, r *http.Request) {
    // common.Encrypt(r.FormValue("plain"))
    // common.Decrypt(r.FormValue("enc"))
    // c := redis_connection()
    // defer c.Close()
    // var key = "KEY"
    // var val = "VALUE"
    // redis_set(key, val, c)
    // s := redis_get(key, c)
    
    id := 12
    common.SetUser(w, r, id)
    common.GetUser(w, r)
    // log.Println("Hello, World")
    fmt.Printf("Path: %s\n", r.URL)
    // str1 := string(r.URL)
    ary2 := strings.Split(r.URL.String(), "/")
    fmt.Printf("ary2: %v\n", ary2)
    // fmt.Println(id)
    // setCookies(w,r)
    // showCookie(w,r)

}

// cookieの設定を行う
func setCookies(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
         Name: "hoge",
         Value: "bar",
    }
    http.SetCookie(w, cookie)

    fmt.Fprintf(w, "Cookieの設定ができたよ")
}

// cookieの取得を行い、htmlを返す
func showCookie(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("hoge")

    if err != nil {
         log.Fatal("Cookie: ", err)
    }

    fmt.Println(cookie)
}

func redis_set(key string, value string, c redis.Conn){
    c.Do("SET", key, value)
  }
  
  func redis_get(key string, c redis.Conn) string{
    s, err := redis.String(c.Do("GET", key))
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    return s
  }
  
  func redis_connection() redis.Conn {
    const IP_PORT = "127.0.0.1:6379"
  
    //redisに接続
    c, err := redis.Dial("tcp", IP_PORT)
    if err != nil {
      panic(err)
    }
    return c
  }