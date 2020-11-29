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

// StockDetail is from minkabu
func StockDetail() {
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
	for rows.Next() {
		if err := rows.Scan(&stockID); err != nil {
			fmt.Println(err)
		}
		res, err := http.Get("https://kabutan.jp/stock/?code=" + stockID)
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
		marketCapitalization := "0"
		profitLast0 := "0"
		profitLast1 := "0"
		profitLast2 := "0"
		doc.Find("#kobetsu_left > table:nth-child(4) > tbody > tr:nth-child(7) > td").Each(func(i2 int, cell *goquery.Selection) {
			fmt.Printf(cell.Text())
			marketCapitalization = strings.Replace(cell.Text(), ",", "", -1)
			marketCapitalization = strings.Replace(marketCapitalization, "億円", "", -1)
		})
		doc.Find("#kobetsu_right > div.gyouseki_block > table > tbody > tr:nth-child(1) > td:nth-child(4)").Each(func(i2 int, cell *goquery.Selection) {
			fmt.Printf(cell.Text())
			profitLast2 = strings.Replace(cell.Text(), ",", "", -1)
			if profitLast2 == "" {
				profitLast2 = "0"
			}
		})
		doc.Find("#kobetsu_right > div.gyouseki_block > table > tbody > tr:nth-child(2) > td:nth-child(4)").Each(func(i2 int, cell *goquery.Selection) {
			fmt.Printf(cell.Text())
			profitLast1 = strings.Replace(cell.Text(), ",", "", -1)
			if profitLast1 == "" {
				profitLast1 = "0"
			}
		})
		doc.Find("#kobetsu_right > div.gyouseki_block > table > tbody > tr:nth-child(3) > td:nth-child(4)").Each(func(i2 int, cell *goquery.Selection) {
			fmt.Printf(cell.Text())
			profitLast0 = strings.Replace(cell.Text(), ",", "", -1)
			if profitLast0 == "" || profitLast0 == "－" {
				profitLast0 = "0"
			}
		})
		fmt.Printf(stockID, stockID, marketCapitalization, profitLast0, profitLast1, profitLast2)
		_, err = db.Exec(`INSERT INTO s_stock (
			stock_id, market_capitalization, profit_last0, profit_last1, profit_last2, invested_at)
		VALUES ($1,$2,$3,$4,$5,$6)`, stockID, marketCapitalization, profitLast0, profitLast1, profitLast2, "2020-11-27")
		if err != nil {
			log.Fatal(err)
		}
	}

	// })
}
