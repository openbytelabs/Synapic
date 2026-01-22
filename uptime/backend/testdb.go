package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./uptime_monitor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM metrics").Scan(&count)
	if err != nil {
		log.Fatal("Metrics tablosu okunamadi:", err)
	}

	fmt.Printf("Toplam kayit sayisi: %d\n\n", count)

	if count > 0 {
		rows, err := db.Query("SELECT url, ping, status, timestamp FROM metrics ORDER BY timestamp DESC LIMIT 10")
		if err != nil {
			log.Fatal("Query hatasi:", err)
		}
		defer rows.Close()

		fmt.Println("Son 10 kayit:")
		fmt.Println("=====================================")
		for rows.Next() {
			var url, status string
			var ping int
			var timestamp time.Time
			rows.Scan(&url, &ping, &status, &timestamp)
			fmt.Printf("URL: %s\nPing: %d ms\nStatus: %s\nTime: %v\n\n", url, ping, status, timestamp)
		}
	} else {
		fmt.Println("Veritabaninda hic kayit yok!")
		fmt.Println("Backend calisiyorsa birkaç saniye bekleyin ve tekrar deneyin.")
	}

	fmt.Println("\nTum URL'ler:")
	rows, _ := db.Query("SELECT DISTINCT url FROM metrics")
	defer rows.Close()
	for rows.Next() {
		var url string
		rows.Scan(&url)
		fmt.Println("-", url)
	}
}
