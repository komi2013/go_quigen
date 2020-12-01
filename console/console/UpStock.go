package console

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq" // this driver for postgres
)

// UpStock how can I run single file??
func UpStock() {
	db, err := sql.Open("postgres", "user=postgres password=sde5tuft dbname=programming sslmode=disable port=5432 host=localhost")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	t := time.Now()
	t = t.AddDate(0, -1, 0)

	query := `select stock_id from s_chart
		where invested_at > $1
		group by stock_id`
	rows, err := db.Query(query, t.Format("2006-1-2"))
	if err != nil {
		fmt.Println(err)
	}
	stockID := ""
	var stockIDs []string
	for rows.Next() {
		if err := rows.Scan(&stockID); err != nil {
			fmt.Println(err)
		}
	}

	res, err := http.Get("https://info.finance.yahoo.co.jp/ranking/?kd=1&tm=d&vl=a&mk=1&p=3")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	stockID = ""
	insert := true
	doc.Find(".rankingTabledata td").Each(func(i2 int, cell *goquery.Selection) {

		if i2%10 == 1 {
			stockID = cell.Text()
		}
		insert = true
		for i := 0; i < len(stockIDs); i++ {
			if stockIDs[i] == stockID {
				insert = false
			}
		}
		if i2%10 == 9 && insert {
			StockHistory(stockID)
		}
	})
}

func DirectInsert() {
	stockIDs := []string{"9273", "3994", "2069"}
	// stockIDs := []string{"4054"}
	for i := 0; i < len(stockIDs); i++ {
		StockHistory(stockIDs[i])
	}
}

// StockHistory insert to s_chart
func StockHistory(stockID string) {
	db, err := sql.Open("postgres", "user=postgres password=sde5tuft dbname=programming sslmode=disable port=5432 host=localhost")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// fmt.Printf(stockID, price, date, ratio, "\n")
	// _, err = db.Exec(`INSERT INTO s_chart (stock_id, price, ratio)
	// 					VALUES ($1,$2,$3)`, stockID, price, ratio)
	// if err != nil {
	// 	log.Print(err)
	// }
	// https://kabutan.jp/stock/kabuka?code=9273
	res, err := http.Get("https://kabutan.jp/stock/kabuka?code=" + stockID)
	// res, err := http.Get("https://stocks.finance.yahoo.co.jp/stocks/history/?code=" + stockID)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	ratio := "0"
	date := ""
	price := ""
	// doc.Find(".boardFin td").Each(func(i3 int, cell2 *goquery.Selection) {
	doc.Find(".stock_kabuka_dwm td").Each(func(i3 int, cell2 *goquery.Selection) {
		fmt.Printf("i3 %#v\n", i3)
		fmt.Printf("cell2.Text() %#v\n", cell2.Text())
		if i3%8 == 0 {
			date = "20" + strings.Replace(cell2.Text(), "/", "-", -1)
			// date = strings.Replace(date, "20", "2020", -1)
			// date = strings.Replace(date, "日", "", -1)
		}
		if i3%8 == 4 {
			price = strings.Replace(cell2.Text(), ",", "", -1)
			fmt.Printf(stockID, price, date, ratio, "\n")
			_, err = db.Exec(`INSERT INTO s_chart (stock_id, price, date, ratio)
								VALUES ($1,$2,$3,$4)`, stockID, price, date, ratio)
			if err != nil {
				log.Print(err)
			}
		}
	})
}
