package app

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// Generate page
func Generate(w http.ResponseWriter, r *http.Request) {

	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT category_id, category_name FROM m_category_name")
	if err != nil {
		// log.Print(query)
		log.Print(r.FormValue("q"))
		log.Print(err)
	}
	category := map[int]string{}
	for rows.Next() {
		categoryID := 0
		categoryName := ""
		if err := rows.Scan(&categoryID, &categoryName); err != nil {
			log.Print(err)
		}
		category[categoryID] = categoryName
	}
	res, _ := json.Marshal(category)
	// fmt.Printf("%#v\n", string(res))
	m := map[string]string{
		"CacheV":   common.CacheV,
		"CSRF":     common.MakeCSRF(w, r),
		"Category": string(res),
	}
	tpl := template.Must(template.ParseFiles("html/generate.html"))
	tpl.Execute(w, m)
}
