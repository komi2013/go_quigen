package app

import (
	"html/template"
	"net/http"

	"../common"

	// "fmt"
	"strings"
)

// Advertisement with iframe
func Advertisement(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("Path: %s\n", r.URL)
	arr := strings.Split(r.URL.String(), "/")
	// fmt.Printf("ary2: %v\n", arr[2])
	tpl := template.Must(template.ParseFiles("html/advertisement/" + arr[2] + ".html"))
	uID := common.GetUser(w, r)
	if uID == 0 {
		return
	}
	m := map[string]string{
		"Date": "Date",
	}
	tpl.Execute(w, m)

}
