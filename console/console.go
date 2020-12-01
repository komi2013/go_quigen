package main

import (
	"os"

	_ "github.com/lib/pq" // this driver for postgres

	"./console"
)

func main() {
	for _, v := range os.Args {
		switch v {
		case "UpStock":
			console.UpStock()
		case "StockDetail":
			console.StockDetail()
		case "DirectInsert":
			console.DirectInsert()
		}
	}
}
