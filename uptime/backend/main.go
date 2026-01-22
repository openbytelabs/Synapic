package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type Monitor struct {
	URL       string
	Name      string
	Language  string
	Tag       string
	CheckJSON bool
	JSONKey   string
}

type Stats struct {
	URL       string  `json:"url"`
	Ping      int     `json:"ping"`
	Status    string  `json:"status"`
	Uptime    float64 `json:"uptime"`
	Timestamp string  `json:"timestamp"`
}

type ErrorLog struct {
	ID           int       `json:"id"`
	Site         string    `json:"site"`
	Status       string    `json:"status"`
	TelegramSent bool      `json:"telegramSent"`
	Message      string    `json:"message"`
	Ping         int       `json:"ping"`
	Timestamp    time.Time `json:"timestamp"`
}

type Metric struct {
	URL       string    `json:"url"`
	Ping      int       `json:"ping"`
	Timestamp time.Time `json:"timestamp"`
}

type UptimeHistory struct {
	URL     string               `json:"url"`
	History map[string]DayUptime `json:"history"`
}

type DayUptime struct {
	Uptime float64 `json:"uptime"`
}

var (
	db               *sql.DB
	monitors         []Monitor
	apiKey           string
	telegramBotToken string
	telegramChatID   string
	lastAlertSent    = make(map[string]time.Time)
	alertCooldown    = 5 * time.Minute
	isFirstCheck     = make(map[string]bool)
	lastDailyReport  time.Time
)

func loadEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Println("UYARI: .env dosyasi bulunamadi, varsayilan degerler kullanilacak")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	log.Println(".env dosyasi yuklendi")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	loadEnv()

	apiKey = getEnv("API_KEY", "your-api-key-here")
	telegramBotToken = getEnv("TELEGRAM_BOT_TOKEN", "")
	telegramChatID = getEnv("TELEGRAM_CHAT_ID", "")

	var err error
	db, err = sql.Open("sqlite3", "./uptime_monitor.db")
	if err != nil {
		log.Fatalf("Veritabani acilamadi: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Veritabanina baglanilamadi: %v", err)
	}

	log.Println("SQLite veritabanina baglandi")

	initDB()

	monitors = []Monitor{
		{URL: "https://synapic.com.tr", Name: "Synapic Main", Language: "TypeScript, Vue.js", Tag: "system", CheckJSON: false, JSONKey: ""},
		{URL: "https://api.synapic.com.tr/api/search?q=test", Name: "Synapic API", Language: "JavaScript", Tag: "system", CheckJSON: true, JSONKey: "results"},
	}

	for _, monitor := range monitors {
		isFirstCheck[monitor.URL] = true
	}

	go monitorLoop()

	router := mux.NewRouter()
	router.HandleFunc("/api/stats", statsHandler).Methods("GET")
	router.HandleFunc("/api/errors", errorsHandler).Methods("GET")
	router.HandleFunc("/api/metrics", metricsHandler).Methods("GET")
	router.HandleFunc("/api/uptime-history", uptimeHistoryHandler).Methods("GET")
	router.HandleFunc("/api/record", recordHandler).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://127.0.0.1:8080"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	port := getEnv("PORT", "3001")
	log.Printf("Server basladi: http://localhost:%s", port)

	if telegramBotToken == "" || telegramChatID == "" {
		log.Println("UYARI: Telegram bilgileri eksik, bildirimler gonderilmeyecek")
	} else {
		log.Println("Telegram bildirimleri aktif")
	}

	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func initDB() {
	createMetrics := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		url TEXT NOT NULL,
		ping INTEGER NOT NULL,
		status TEXT NOT NULL,
		timestamp DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_url ON metrics(url);
	CREATE INDEX IF NOT EXISTS idx_timestamp ON metrics(timestamp);
	`

	createErrors := `
	CREATE TABLE IF NOT EXISTS error_logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		site TEXT NOT NULL,
		status TEXT NOT NULL,
		telegram_sent BOOLEAN NOT NULL DEFAULT 0,
		message TEXT,
		ping INTEGER DEFAULT 0,
		timestamp DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_site ON error_logs(site);
	CREATE INDEX IF NOT EXISTS idx_error_timestamp ON error_logs(timestamp);
	`

	_, err := db.Exec(createMetrics)
	if err != nil {
		log.Fatal("Metrics tablosu olusturulamadi:", err)
	}

	_, err = db.Exec(createErrors)
	if err != nil {
		log.Fatal("Error_logs tablosu olusturulamadi:", err)
	}

	log.Println("Veritabani tablolari hazir")
}

func monitorLoop() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	log.Println("Monitoring basladi (5 saniye araliklarla)")

	for _, monitor := range monitors {
		go checkSite(monitor)
	}

	go dailyReportScheduler()

	for range ticker.C {
		for _, monitor := range monitors {
			go checkSite(monitor)
		}
	}
}

func checkSite(monitor Monitor) {
	start := time.Now()

	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	req, err := http.NewRequest("GET", monitor.URL, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "Synapic-Monitor/1.0")

	resp, err := client.Do(req)
	duration := time.Since(start).Milliseconds()

	status := "online"
	var errorMsg string

	if err != nil {
		status = "offline"
		errorMsg = err.Error()

		if !isFirstCheck[monitor.URL] {
			logError(monitor.URL, status, errorMsg, int(duration))
			sendTelegramAlert(monitor, errorMsg, int(duration))
		}
	} else if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		status = "offline"
		errorMsg = fmt.Sprintf("HTTP %d", resp.StatusCode)

		if !isFirstCheck[monitor.URL] {
			logError(monitor.URL, status, errorMsg, int(duration))
			sendTelegramAlert(monitor, errorMsg, int(duration))
		}
	} else if monitor.CheckJSON && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		body := make([]byte, 10240)
		n, _ := resp.Body.Read(body)
		bodyStr := string(body[:n])

		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(bodyStr), &jsonData); err != nil {
			status = "offline"
			errorMsg = "JSON parse hatasi"

			if !isFirstCheck[monitor.URL] {
				logError(monitor.URL, status, errorMsg, int(duration))
				sendTelegramAlert(monitor, errorMsg, int(duration))
			}
		} else {
			if results, exists := jsonData[monitor.JSONKey]; !exists || results == nil {
				status = "offline"
				errorMsg = fmt.Sprintf("JSON'da '%s' anahtari bulunamadi veya null", monitor.JSONKey)

				if !isFirstCheck[monitor.URL] {
					logError(monitor.URL, status, errorMsg, int(duration))
					sendTelegramAlert(monitor, errorMsg, int(duration))
				}
			} else {
				if resultArray, ok := results.([]interface{}); ok && len(resultArray) == 0 {
					status = "offline"
					errorMsg = fmt.Sprintf("JSON'da '%s' bos array", monitor.JSONKey)

					if !isFirstCheck[monitor.URL] {
						logError(monitor.URL, status, errorMsg, int(duration))
						sendTelegramAlert(monitor, errorMsg, int(duration))
					}
				}
			}
		}
	}

	if resp != nil {
		resp.Body.Close()
	}

	if isFirstCheck[monitor.URL] {
		isFirstCheck[monitor.URL] = false
	}

	recordMetric(monitor.URL, int(duration), status)
}

func recordMetric(url string, ping int, status string) {
	query := "INSERT INTO metrics (url, ping, status, timestamp) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, url, ping, status, time.Now())
	if err != nil {
		log.Printf("HATA: Metric kaydedilemedi: %v", err)
	}
}

func logError(site, status, message string, ping int) {
	_, err := db.Exec("INSERT INTO error_logs (site, status, telegram_sent, message, ping, timestamp) VALUES (?, ?, ?, ?, ?, ?)",
		site, status, false, message, ping, time.Now())
	if err != nil {
		log.Printf("Hata kaydedilemedi: %v", err)
	}
}

func sendTelegramAlert(monitor Monitor, errorMsg string, ping int) {
	if telegramBotToken == "" || telegramChatID == "" {
		return
	}

	lastSent, exists := lastAlertSent[monitor.URL]
	if exists && time.Since(lastSent) < alertCooldown {
		return
	}

	message := fmt.Sprintf(
		"[SUNUCU COKTU]\n\n"+
			"Site: %s\n"+
			"URL: %s\n"+
			"Durum: OFFLINE\n"+
			"Ping: %d ms\n"+
			"Hata: %s\n"+
			"Zaman: %s\n"+
			"Tag: %s",
		monitor.Name,
		monitor.URL,
		ping,
		errorMsg,
		time.Now().Format("02.01.2006 15:04:05"),
		monitor.Tag,
	)

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)

	payload := map[string]interface{}{
		"chat_id": telegramChatID,
		"text":    message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		updateTelegramStatus(monitor.URL, false)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		lastAlertSent[monitor.URL] = time.Now()
		updateTelegramStatus(monitor.URL, true)
		log.Printf("Telegram bildirimi gonderildi: %s - %s", monitor.Name, errorMsg)
	} else {
		updateTelegramStatus(monitor.URL, false)
	}
}

func updateTelegramStatus(site string, sent bool) {
	_, err := db.Exec(
		"UPDATE error_logs SET telegram_sent = ? WHERE site = ? AND id = (SELECT id FROM error_logs WHERE site = ? ORDER BY timestamp DESC LIMIT 1)",
		sent, site, site)
	if err != nil {
		log.Printf("Telegram durumu guncellenemedi: %v", err)
	}
}

func dailyReportScheduler() {
	for {
		now := time.Now()
		nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		duration := nextMidnight.Sub(now)

		log.Printf("Gunluk rapor %v sonra gonderilecek", duration)

		time.Sleep(duration)
		sendDailyReport()
	}
}

func sendDailyReport() {
	if telegramBotToken == "" || telegramChatID == "" {
		return
	}

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	startOfYesterday := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())
	endOfYesterday := startOfYesterday.Add(24 * time.Hour)

	var report strings.Builder
	report.WriteString(fmt.Sprintf("📊 GUNLUK DURUM RAPORU\n"))
	report.WriteString(fmt.Sprintf("📅 Tarih: %s\n\n", yesterday.Format("02.01.2006")))

	for _, monitor := range monitors {
		report.WriteString(fmt.Sprintf("🔹 %s\n", monitor.Name))
		report.WriteString(fmt.Sprintf("   URL: %s\n", monitor.URL))

		var total, online, totalPing int
		var minPing, maxPing int = 999999, 0

		rows, err := db.Query(
			"SELECT status, ping FROM metrics WHERE url = ? AND timestamp >= ? AND timestamp < ?",
			monitor.URL, startOfYesterday, endOfYesterday)

		if err == nil {
			for rows.Next() {
				var status string
				var ping int
				rows.Scan(&status, &ping)
				total++
				totalPing += ping
				if status == "online" {
					online++
				}
				if ping < minPing {
					minPing = ping
				}
				if ping > maxPing {
					maxPing = ping
				}
			}
			rows.Close()
		}

		if total > 0 {
			uptime := (float64(online) / float64(total)) * 100
			avgPing := totalPing / total

			var uptimeEmoji string
			if uptime >= 99.9 {
				uptimeEmoji = "✅"
			} else if uptime >= 95 {
				uptimeEmoji = "⚠️"
			} else {
				uptimeEmoji = "❌"
			}

			report.WriteString(fmt.Sprintf("   %s Uptime: %.2f%%\n", uptimeEmoji, uptime))
			report.WriteString(fmt.Sprintf("   📈 Ortalama Ping: %d ms\n", avgPing))
			report.WriteString(fmt.Sprintf("   ⬇️ Min Ping: %d ms\n", minPing))
			report.WriteString(fmt.Sprintf("   ⬆️ Max Ping: %d ms\n", maxPing))
			report.WriteString(fmt.Sprintf("   🔢 Toplam Kontrol: %d\n", total))
			report.WriteString(fmt.Sprintf("   ✓ Basarili: %d\n", online))
			report.WriteString(fmt.Sprintf("   ✗ Basarisiz: %d\n", total-online))
		} else {
			report.WriteString("   ⚠️ Veri yok\n")
		}

		var errorCount int
		db.QueryRow(
			"SELECT COUNT(*) FROM error_logs WHERE site = ? AND timestamp >= ? AND timestamp < ?",
			monitor.URL, startOfYesterday, endOfYesterday).Scan(&errorCount)

		if errorCount > 0 {
			report.WriteString(fmt.Sprintf("   🚨 Hata Sayisi: %d\n", errorCount))
		}

		report.WriteString("\n")
	}

	var totalErrors int
	db.QueryRow(
		"SELECT COUNT(*) FROM error_logs WHERE timestamp >= ? AND timestamp < ?",
		startOfYesterday, endOfYesterday).Scan(&totalErrors)

	report.WriteString(fmt.Sprintf("📌 OZET\n"))
	report.WriteString(fmt.Sprintf("Toplam Hata: %d\n", totalErrors))
	report.WriteString(fmt.Sprintf("Rapor Saati: %s\n", now.Format("02.01.2006 15:04:05")))

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramBotToken)

	payload := map[string]interface{}{
		"chat_id": telegramChatID,
		"text":    report.String(),
	}

	jsonData, _ := json.Marshal(payload)

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			log.Println("Gunluk rapor Telegram'a gonderildi")
			lastDailyReport = now
		}
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("apiKey") != apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	stats := make([]Stats, 0)

	for _, monitor := range monitors {
		var ping int
		var status string
		err := db.QueryRow("SELECT ping, status FROM metrics WHERE url = ? ORDER BY timestamp DESC LIMIT 1",
			monitor.URL).Scan(&ping, &status)

		uptime := calculateUptime(monitor.URL)

		stat := Stats{
			URL:       monitor.URL,
			Ping:      ping,
			Status:    status,
			Uptime:    uptime,
			Timestamp: time.Now().Format(time.RFC3339),
		}

		if err != nil {
			stat.Status = "unknown"
			stat.Ping = 0
		}

		stats = append(stats, stat)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

func calculateUptime(url string) float64 {
	var total, online int

	rows, err := db.Query("SELECT status FROM metrics WHERE url = ? AND timestamp > datetime('now', '-24 hours')", url)
	if err != nil {
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		rows.Scan(&status)
		total++
		if status == "online" {
			online++
		}
	}

	if total == 0 {
		return 100
	}

	return (float64(online) / float64(total)) * 100
}

func calculateDailyUptime(url string, date time.Time) float64 {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var total, online int

	rows, err := db.Query(
		"SELECT status FROM metrics WHERE url = ? AND timestamp >= ? AND timestamp < ?",
		url, startOfDay, endOfDay)
	if err != nil {
		return 100
	}
	defer rows.Close()

	for rows.Next() {
		var status string
		rows.Scan(&status)
		total++
		if status == "online" {
			online++
		}
	}

	if total == 0 {
		return 100
	}

	return (float64(online) / float64(total)) * 100
}

func uptimeHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("apiKey") != apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	histories := make([]UptimeHistory, 0)

	for _, monitor := range monitors {
		history := make(map[string]DayUptime)
		now := time.Now()

		for i := 89; i >= 0; i-- {
			date := now.AddDate(0, 0, -i)
			dateStr := date.Format("2006-01-02")
			uptime := calculateDailyUptime(monitor.URL, date)

			history[dateStr] = DayUptime{
				Uptime: uptime,
			}
		}

		histories = append(histories, UptimeHistory{
			URL:     monitor.URL,
			History: history,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(histories)
}

func errorsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("apiKey") != apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	rows, err := db.Query("SELECT id, site, status, telegram_sent, message, ping, timestamp FROM error_logs ORDER BY timestamp DESC LIMIT 100")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	logs := make([]ErrorLog, 0)
	for rows.Next() {
		var log ErrorLog
		err := rows.Scan(&log.ID, &log.Site, &log.Status, &log.TelegramSent, &log.Message, &log.Ping, &log.Timestamp)
		if err != nil {
			continue
		}
		logs = append(logs, log)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("apiKey") != apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	timeRange := r.URL.Query().Get("range")
	site := r.URL.Query().Get("site")

	var timeFilter string

	switch timeRange {
	case "minute":
		timeFilter = "datetime('now', '-5 minutes')"
	case "hour":
		timeFilter = "datetime('now', '-1 hour')"
	case "day":
		timeFilter = "datetime('now', '-24 hours')"
	default:
		timeFilter = "datetime('now', '-5 minutes')"
	}

	query := fmt.Sprintf("SELECT url, ping, timestamp FROM metrics WHERE timestamp > %s", timeFilter)

	if site != "" && site != "all" {
		query += fmt.Sprintf(" AND url = '%s'", site)
	}

	query += " ORDER BY timestamp ASC LIMIT 500"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	metrics := make([]Metric, 0)
	for rows.Next() {
		var metric Metric
		err := rows.Scan(&metric.URL, &metric.Ping, &metric.Timestamp)
		if err != nil {
			continue
		}
		metrics = append(metrics, metric)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(metrics)
}

func recordHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("apiKey") != apiKey {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var metric Metric
	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := "online"
	if metric.Ping > 5000 {
		status = "offline"
	}

	recordMetric(metric.URL, metric.Ping, status)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
