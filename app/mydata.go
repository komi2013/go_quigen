package app

import (
    "net/http"
    "html/template"
    // "github.com/garyburd/redigo/redis"
    // "fmt"
    // "os"
    // "log"
    // "time"
    "../common"
    // "net/url"
    // "strings"
)

func Mydata(w http.ResponseWriter, r *http.Request) {
    type View struct {
      CacheV string
      CSRF string
      Myphoto string
      EtoColor string
      DateLanguage string
    }
    var view View
    view.CacheV = common.CACHE_V
    view.CSRF = common.MakeCSRF(w,r)
    view.DateLanguage = common.DATE_LANGUAGE

    u_id := common.GetUser(w,r)
    eto := common.Eto(u_id)
    view.Myphoto = eto[0]
    view.EtoColor = eto[1]

    tpl := template.Must(template.ParseFiles("html/mydata.html"))
    tpl.Execute(w, view)

}
