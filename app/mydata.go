package app

import (
	"html/template"
	"net/http"

	"../common"
)

// Mydata page
func Mydata(w http.ResponseWriter, r *http.Request) {
	type View struct {
		CacheV       string
		CSRF         string
		Myphoto      string
		EtoColor     string
		DateLanguage string
	}
	var view View
	view.CacheV = common.CacheV
	view.CSRF = common.MakeCSRF(w, r)
	view.DateLanguage = common.DateLanguage

	uID := common.GetUser(w, r)
	eto := common.Eto(uID)
	view.Myphoto = eto[0]
	view.EtoColor = eto[1]

	tpl := template.Must(template.ParseFiles("html/mydata.html"))
	tpl.Execute(w, view)

}
