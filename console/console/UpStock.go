package console

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	t = t.AddDate(0, 0, -20)

	query := `select stock_id from s_chart
		where chart_id = 1269
		`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	stockID := ""
	// var stockIDs []string
	for rows.Next() {
		if err := rows.Scan(&stockID); err != nil {
			fmt.Println(err)
		}

		// stockIDs = append(stockIDs, stockID)
	}
	fmt.Println(stockID)

	// urls := []string{"https://info.finance.yahoo.co.jp/ranking/?kd=1&mk=1&tm=d&vl=a", "https://info.finance.yahoo.co.jp/ranking/?kd=1&tm=d&vl=a&mk=1&p=2", "https://info.finance.yahoo.co.jp/ranking/?kd=1&tm=d&vl=a&mk=1&p=3"}
	// for i := 0; i < len(urls); i++ {
	// 	Scraping(urls[i], stockIDs, db)
	// }
	// query = `select stock_id from s_chart
	// 	group by stock_id
	// 	having count(*) < 2`
	// rows, err = db.Query(query)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for rows.Next() {
	// 	if err := rows.Scan(&stockID); err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	yahooHistory(stockID)
	// }
	// StockDetail()
}

// Scraping focus on scraping
func Scraping(url string, stockIDs []string, db *sql.DB) {
	res, err := http.Get(url)
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
	stockID := ""
	price := ""
	ratio := ""
	insert := true
	doc.Find(".rankingTabledata td").Each(func(i2 int, cell *goquery.Selection) {
		if i2%10 == 1 {
			stockID = cell.Text()

		}
		if i2%10 == 5 {
			price = strings.Replace(cell.Text(), ",", "", -1)
		}
		if i2%10 == 6 {
			ratio = strings.Replace(cell.Text(), "%", "", -1)
		}
		insert = true
		for i := 0; i < len(stockIDs); i++ {
			if stockIDs[i] == stockID {
				insert = false
			}
		}
		if i2%10 == 9 && insert {
			fmt.Printf(stockID, price, ratio, "\n")
			_, err = db.Exec(`INSERT INTO s_chart (stock_id, price, ratio)
				VALUES ($1,$2,$3)`, stockID, price, ratio)
			if err != nil {
				log.Print(err)
			}
			StockHistory(stockID)
		}
	})
}

// DirectInsert from manual data
func DirectInsert() {
	stockIDs := []string{"6076"}
	// stockIDs := []string{"4481"}
	for i := 0; i < len(stockIDs); i++ {
		// StockHistory(stockIDs[i])
		yahooHistory(stockIDs[i])
	}
}

// StockHistory insert to s_chart
func StockHistory(stockID string) {
	db, err := sql.Open("postgres", "user=postgres password=sde5tuft dbname=programming sslmode=disable port=5432 host=localhost")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	res, err := http.Get("https://kabutan.jp/stock/kabuka?code=" + stockID)
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
	for i := 1; i < 30; i++ {
		s := strconv.Itoa(i)
		doc.Find("#stock_kabuka_table > table.stock_kabuka_dwm > tbody > tr:nth-child(" + s + ") > th").Each(func(i3 int, v *goquery.Selection) {
			date = "20" + strings.Replace(v.Text(), "/", "-", -1)
		})
		doc.Find("#stock_kabuka_table > table.stock_kabuka_dwm > tbody > tr:nth-child(" + s + ") > td:nth-child(5)").Each(func(i3 int, v *goquery.Selection) {
			price = strings.Replace(v.Text(), ",", "", -1)
		})
		doc.Find("#stock_kabuka_table > table.stock_kabuka_dwm > tbody > tr:nth-child(" + s + ") > td:nth-child(7) > span").Each(func(i3 int, v *goquery.Selection) {
			ratio = v.Text()
		})
		fmt.Printf(stockID, price, date, ratio, "\n")
		_, err = db.Exec(`INSERT INTO s_chart (stock_id, price, date, ratio)
					VALUES ($1,$2,$3,$4)`, stockID, price, date, ratio)
		if err != nil {
			log.Print(err)
		}
	}
}
func yahooHistory(stockID string) {
	db, err := sql.Open("postgres", "user=postgres password=sde5tuft dbname=programming sslmode=disable port=5432 host=localhost")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	res, err := http.Get("https://stocks.finance.yahoo.co.jp/stocks/history/?code=" + stockID)
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
	date := ""
	price := ""
	doc.Find(".boardFin td").Each(func(i3 int, cell2 *goquery.Selection) {
		// fmt.Printf("td %s\n", cell2.Text())

		if i3%7 == 0 {
			date = strings.Replace(cell2.Text(), "年", "-", -1)
			date = strings.Replace(date, "月", "-", -1)
			date = strings.Replace(date, "日", "", -1)
		}
		if i3%7 == 6 {

			price = strings.Replace(cell2.Text(), ",", "", -1)
			fmt.Printf(stockID, price, date, "\n")
			_, err := db.Exec(`INSERT INTO s_chart (stock_id, price, date)
								VALUES ($1,$2,$3)`, stockID, price, date)
			if err != nil {
				log.Print(err)
			}
		}
	})
}
