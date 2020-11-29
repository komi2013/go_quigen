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
	// fmt.Println(t.Format("2006-1-2"))

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
		// stockIDs = append(stockIDs, stockID)
	}

	res, err := http.Get("https://info.finance.yahoo.co.jp/ranking/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	stockID = ""
	ratio := "0"
	date := ""
	price := ""
	insert := true
	doc.Find(".rankingTabledata td").Each(func(i2 int, cell *goquery.Selection) {
		if i2%10 == 1 {
			stockID = cell.Text()
		}
		if i2%10 == 5 {
			price = strings.Replace(cell.Text(), ",", "", -1)
		}
		if i2%10 == 6 {
			ratio = cell.Text()
			ratio = strings.Replace(ratio, "+", "", -1)
			ratio = strings.Replace(ratio, "%", "", -1)
		}
		fmt.Printf("stockID: %v\n", stockID)
		insert = true
		for i := 0; i < len(stockIDs); i++ {
			if stockIDs[i] == stockID {
				insert = false
			}
		}
		if i2%10 == 9 && insert {
			fmt.Printf(stockID, price, date, ratio)
			_, err = db.Exec(`INSERT INTO s_chart (stock_id, price, date, ratio)
								VALUES ($1,$2,$3,$4)`, stockID, price, "2020-11-27", ratio)
			if err != nil {
				log.Print(err)
			}
			// res, err = http.Get("https://stocks.finance.yahoo.co.jp/stocks/history/?code=" + stockID + ".T")
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// defer res.Body.Close()
			// if res.StatusCode != 200 {
			// 	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
			// }
			// doc, err := goquery.NewDocumentFromReader(res.Body)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// doc.Find(".boardFin td").Each(func(i3 int, cell2 *goquery.Selection) {
			// 	// fmt.Printf("td %s\n", cell2.Text())

			// 	if i3%7 == 0 {
			// 		date = strings.Replace(cell2.Text(), "年", "-", -1)
			// 		date = strings.Replace(date, "月", "-", -1)
			// 		date = strings.Replace(date, "日", "", -1)
			// 	}
			// 	if i3%7 == 6 {

			// 		price = strings.Replace(cell2.Text(), ",", "", -1)
			// 		fmt.Printf(stockID, price, date, ratio)
			// 		_, err = db.Exec(`INSERT INTO s_chart (stock_id, price, date, up_ratio)
			// 							VALUES ($1,$2,$3,$4)`, stockID, price, date, ratio)
			// 		if err != nil {
			// 			log.Print(err)
			// 		}
			// 	}
			// })

		}
		// fmt.Printf("rank %s, stockID %s\n", i2, stockID)
	})
	// })
}
