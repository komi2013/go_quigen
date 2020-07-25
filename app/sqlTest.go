package app

import (
  "net/http"
  "database/sql"
  "fmt"
  "github.com/go-gorp/gorp"
  _ "github.com/lib/pq"
  "log"
  "../common"
)

type seq struct {
  Nextval    uint32 
  Usr_name  string `db:"usr_name"` // FirstName に name というカラムを紐付ける
}

func SqlTest(w http.ResponseWriter, r *http.Request) {



	connStr := "user=exam_8099 password=Zk1CGsBK dbname=go_english sslmode=disable port=5432 host=194.135.95.163"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	question_id := 3
	rows, err := db.Query("select question_id,question_txt from t_question WHERE question_id = $1 AND usr_id = $2", question_id,0)
  if err != nil {
      panic(err)
  }
  var question []common.TQuestion
  for rows.Next() {
      q := common.TQuestion{}
      if err := rows.Scan(&q.QuestionID,&q.QuestionTxt); err != nil {
          log.Fatal(err)
      }

      question = append(question,q)
  }
  fmt.Printf("%v\n", question)
  // _, err = db.Exec(`DELETE FROM t_question WHERE question_id = $1`, 4)
  // if err != nil {
  //     panic(err)
  // }
  // ins, err := db.Prepare("INSERT INTO t_question(question_id,question_txt) VALUES(?,?)")
  // if err != nil {
  //     log.Fatal(err)
  // }
  // ins.Exec(5,"golang+001@gmail.com")
  // dbmap := initDb()
  // defer dbmap.Db.Close()
  // rows, err := dbmap.Select(seq{}, "select nextval('t_question_question_id_seq')")
  // rows, err = dbmap.Select(table.TQuestion{}, "select question_id from t_question")
  // if err != nil {
  //   log.Fatalln(err)
  // }
  // var test string
  // for i, r := range rows {
  //     d := r.(*table.TQuestion)
  //     fmt.Printf("[%d] id: %d \n", i, d.QuestionID)
  //     // test = row.usr_name
  //     // fmt.Printf("%v\n", row.usr_name)
  // }
  // fmt.Printf("%v\n", test)
}

func initDb() *gorp.DbMap {
  db, _ := sql.Open("postgres",
      "user=exam_8099 password=Zk1CGsBK dbname=go_english sslmode=disable port=5432 host=194.135.95.163")
  // db, _ := sql.Open("postgres",
  // "exam_8099:Zk1CGsBK@tcp(194.135.95.163:5432)/english")
  // db, _ := sql.Open("postgres",
  // "tcp:194.135.95.163:5432*english/postgres/sde5tuft")
  // "tcp:localhost:3306*mydb/myuser/mypassword"
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
  // dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Usr_id")
  return dbmap
}