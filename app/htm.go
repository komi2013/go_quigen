package app

import (
	"html/template"
	"net/http"
	"strings"

	"../common"
)

// Htm can make many html pages
func Htm(w http.ResponseWriter, r *http.Request) {
	type View struct {
		CacheV   string
		CSRF     string
		Myphoto  string
		EtoColor string
	}
	var view View
	view.CacheV = common.CacheV
	view.CSRF = common.MakeCSRF(w, r)

	uID := common.GetUser(w, r)
	eto := common.Eto(uID)
	view.Myphoto = eto[0]
	view.EtoColor = eto[1]
	uri := strings.Split(r.URL.String(), "/")
	tpl := template.Must(template.ParseFiles("html/htm/" + uri[2] + ".html"))
	tpl.Execute(w, view)

}
