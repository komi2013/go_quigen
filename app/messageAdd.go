package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"../common"
	_ "github.com/lib/pq" // this driver for postgres
)

// MessageAdd insert t_message
func MessageAdd(w http.ResponseWriter, r *http.Request) {
	uID := common.GetUser(w, r)
	ok := common.CheckCSRF(r, r.FormValue("csrf"))
	if !ok {
		fmt.Fprint(w, `["2","CSRF Error"]`)
		return
	}
	db, err := sql.Open("postgres", common.DbConnect)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	_, err = db.Exec(`INSERT INTO t_message(
        sender,receiver,message_content,created_at,sender_img
        ) VALUES($1,$2,$3,$4,$5)`,
		uID,
		r.FormValue("receiver"),
		r.FormValue("message_content"),
		time.Now().Format("2006-01-02 15:04:05"),
		r.FormValue("sender_img"))
	if err != nil {
		log.Print(err)
	}
	fmt.Fprint(w, `{"Status":"1"}`)

}
