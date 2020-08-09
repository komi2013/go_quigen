package app

import (
    "net/http"
    "html/template"
    // "github.com/garyburd/redigo/redis"
    "fmt"
    // "os"
    // "log"
    // "time"
    "../common"
    // "net/url"
    "strings"
)

func Htm(w http.ResponseWriter, r *http.Request) {
    type View struct {
      CacheV string
      CSRF string
      Myphoto string
      EtoColor string
    }
    var view View
    view.CacheV = common.CACHE_V
    view.CSRF = common.MakeCSRF(w,r)

    u_id := common.GetUser(w,r)
    eto := common.Eto(u_id)
    view.Myphoto = eto[0]
    view.EtoColor = eto[1]
    fmt.Printf("eto %#v\n", eto)
    uri := strings.Split(r.URL.String(), "/")
    tpl := template.Must(template.ParseFiles("html/htm/"+uri[2]+".html"))
    tpl.Execute(w, view)

}
