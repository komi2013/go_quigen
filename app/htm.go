package app

import (
    "net/http"
    "html/template"
    // "github.com/garyburd/redigo/redis"
    "fmt"
    // "os"
    // "log"
    // "time"
    // "../common"
    // "net/url"
    "strings"
)

func Htm(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Path: %s\n", r.URL)
    arr := strings.Split(r.URL.String(), "/")
    fmt.Printf("ary2: %v\n", arr[2])
    tpl := template.Must(template.ParseFiles("ad/ad_menu.html"))
    m := map[string]string{
      "Date": "Date",
      "Time": "Time",
      "Testt": "near",
    }
    tpl.Execute(w, m)

}
