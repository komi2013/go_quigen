package app

import (
    // "fmt"
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    // "time"
    "../common"
    // "./app"
    // "crypto/aes"
)

func Generate(w http.ResponseWriter, r *http.Request) {

	connStr := common.DB_CONNECT
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
    }
    defer db.Close()
    var cat string
    cookie, err := r.Cookie("cat")
    if err != nil {
		cat = "0"
    } else {
        cat = cookie.Value
    }
    
    query := "SELECT category_name FROM m_category_name WHERE category_id = $1"
    rows, err := db.Query(query, cat)
    if err != nil {
        log.Print(query)
        log.Print(r.FormValue("q"))
        log.Print(err)
    }
    category_name := ""
    for rows.Next() {
        if err := rows.Scan(&category_name); err != nil {
            log.Print(err)
        }
    }
    m := map[string]string{
        "CacheV": common.CACHE_V,
        "CSRF": common.MakeCSRF(w,r),
        "CategoryName": category_name,
    }
    tpl := template.Must(template.ParseFiles("html/generate.html"))
    tpl.Execute(w, m)
}