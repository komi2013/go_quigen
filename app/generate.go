package app

import (
    "fmt"
    "html/template"
    // "log"
    "net/http"
    "time"
    "../common"
    // "./app"
    // "crypto/aes"
)

func Generate(w http.ResponseWriter, r *http.Request) {


    
    m := map[string]string{
        "cache_v": common.CACHE_V,
        "csrf": common.MakeCSRF(w,r),
        "Time": time.Now().Format("15:04:05"),
        "Testt": "near",
    }
    cookie, err := r.Cookie("xsssr")
    value := ""
    if err == nil {
        value = cookie.Value
	}
    // fmt.Println("koutei")
    fmt.Printf("%#v\n", value)
    tpl := template.Must(template.ParseFiles("html/generate.html"))
    tpl.Execute(w, m)
}