package app

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// Search page
func Search(w http.ResponseWriter, r *http.Request) {
	connStr := common.DbConnect
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT category_id FROM m_category_name WHERE category_name = $1", r.FormValue("tag"))
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		var categoryID int
		if err := rows.Scan(&categoryID); err != nil {
			log.Print(err)
		}
		http.Redirect(w, r, "/category/1/"+strconv.Itoa(categoryID)+"/", 301)
		return
	}
}
