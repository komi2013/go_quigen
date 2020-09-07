package app

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    // "log"
    "net/http"
    "../common"
    "time"
    "net/smtp"
    "encoding/base64"
)

func CommentAdd(w http.ResponseWriter, r *http.Request) {

    ok := common.CheckCSRF(r,r.FormValue("csrf"))
    if !ok {
        fmt.Fprint(w, `["2","CSRF Error"]`)
        return
    }
	db, err := sql.Open("postgres", common.DB_CONNECT)
	if err != nil {
		fmt.Println(err)
    }
	defer db.Close()
	u_id := common.GetUser(w,r)
    _, err = db.Exec(`INSERT INTO t_comment(
        comment_txt,usr_id,question_id,created_at,usr_img
        ) VALUES($1,$2,$3,$4,$5)`, 
        r.FormValue("comment_txt"),
        u_id,
        r.FormValue("question_id"),
        time.Now().Format("2006-01-02 15:04:05"),
        r.FormValue("usr_img") )
    if err != nil {
        fmt.Println(err)
    }

	// Connect to the remote SMTP server.
	c, err := smtp.Dial("localhost:25")
	if err != nil {
		fmt.Println(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("sender@quigen.info"); err != nil {
		fmt.Println(err)
	}
	if err := c.Rcpt("seijirok@infoseek.jp"); err != nil {
		fmt.Println(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		fmt.Println(err)
    }
    //" + strings.Join(to, ",") + "
    msg := "To: seiijrok@infoseek.jp \r\n" +
		"From: sender@quigen.info \r\n" +
		"Subject: you got comment \r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(r.FormValue("comment_txt")))
    // _, err = fmt.Fprintf(wc, "This is the email body")
    _, err = wc.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	err = wc.Close()
	if err != nil {
		fmt.Println(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		fmt.Println(err)
	}

    fmt.Fprint(w, `{"Status":"1"}`)

}