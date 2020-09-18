package app

import (
    // "fmt"
    "html/template"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "net/http"
    "encoding/json"
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
    rows, err := db.Query("SELECT category_id, category_name FROM m_category_name")
    if err != nil {
        // log.Print(query)
        log.Print(r.FormValue("q"))
        log.Print(err)
    }
    category := map[int]string{}
    for rows.Next() {
        category_id := 0
        category_name := ""
        if err := rows.Scan(&category_id,&category_name); err != nil {
            log.Print(err)
        }
        category[category_id] = category_name
    }
    res, _ := json.Marshal(category)
    // fmt.Printf("%#v\n", string(res))
    m := map[string]string{
        "CacheV": common.CACHE_V,
        "CSRF": common.MakeCSRF(w,r),
        "Category": string(res),
    }
    tpl := template.Must(template.ParseFiles("html/generate.html"))
    tpl.Execute(w, m)
}