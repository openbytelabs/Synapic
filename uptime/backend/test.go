package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func loadEnvTest() map[string]string {
	env := make(map[string]string)
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal("env dosyasi acilamadi:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			env[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return env
}

func main() {
	env := loadEnvTest()

	token := env["TELEGRAM_BOT_TOKEN"]
	chatID := env["TELEGRAM_CHAT_ID"]

	fmt.Printf("Token: %s (uzunluk: %d)\n", token, len(token))
	fmt.Printf("Chat ID: %s\n\n", chatID)

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    "Test mesaji - Sistem calisiyor",
	}

	jsonData, _ := json.Marshal(payload)

	fmt.Println("Gonderilen JSON:")
	fmt.Println(string(jsonData))
	fmt.Println()

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Istek gonderilemedi:", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("HTTP Durum Kodu: %d\n", resp.StatusCode)
	fmt.Println("Yanit:")
	fmt.Println(string(body))

	if resp.StatusCode == 200 {
		fmt.Println("\nBasarili! Telegram mesaji gonderildi.")
	} else {
		fmt.Println("\nHata! Mesaj gonderilemedi.")
	}
}
