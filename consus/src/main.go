package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/blugelabs/bluge"
)

type SearchResult struct {
	URL         string  `json:"url"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Score       float64 `json:"score"`
}

type SearchResponse struct {
	Results    []SearchResult `json:"results"`
	Total      int            `json:"total"`
	TotalUsed  int            `json:"total_used,omitempty"`
	Remaining  int            `json:"remaining,omitempty"`
	DailyLimit int            `json:"daily_limit,omitempty"`
}

type IndexStatsResponse struct {
	IndexPath string  `json:"index_path"`
	SizeBytes int64   `json:"size_bytes"`
	SizeMB    float64 `json:"size_mb"`
	SizeGB    float64 `json:"size_gb"`
}

type User struct {
	GoogleID    string `json:"google_id"`
	Email       string `json:"email"`
	APIKey      string `json:"apikey"`
	SearchCount int    `json:"search_count"`
	ResetDate   string `json:"reset_date"`
	DailyLimit  int    `json:"daily_limit"`
	IsMaster    bool   `json:"is_master"`
}

type UserStore struct {
	mu        sync.RWMutex
	users     map[string]*User
	masterKey string
}

type AuthResponse struct {
	Success bool   `json:"success"`
	APIKey  string `json:"apikey,omitempty"`
	Message string `json:"message,omitempty"`
}

type StatsResponse struct {
	SearchesToday int `json:"searches_today"`
	DailyLimit    int `json:"daily_limit"`
	Remaining     int `json:"remaining"`
}

var userStore = &UserStore{
	users: make(map[string]*User),
}

var serverIP string

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowedDomain := "https://synapic.com.tr"
		serverOrigin := fmt.Sprintf("http://%s", serverIP)
		serverOriginWithPort := fmt.Sprintf("http://%s:3000", serverIP)
		isAllowed := origin == allowedDomain ||
			origin == serverOrigin ||
			origin == serverOriginWithPort ||
			strings.Contains(origin, "localhost") ||
			strings.Contains(origin, "127.0.0.1")

		if isAllowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else if origin != "" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "CORS Error: Origin %s is not allowed", origin)
			return
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func getServerIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("Error getting network interfaces: %v", err)
		return ""
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getClientIP(r *http.Request) string {
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return xri
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func apiIPRestriction(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)

		if clientIP == "127.0.0.1" || clientIP == "::1" || clientIP == serverIP {
			next(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Access denied: This endpoint is restricted to server IP only",
		})
	}
}

func generateAPIKey() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return "synapic_" + hex.EncodeToString(bytes)
}

func generateMasterKey() string {
	bytes := make([]byte, 24)
	rand.Read(bytes)
	return "synapic_master_" + hex.EncodeToString(bytes)
}

func getTodayDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func getNextMidnight() time.Time {
	now := time.Now()
	nextMidnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return nextMidnight
}

func startMidnightResetScheduler() {
	go func() {
		for {
			nextMidnight := getNextMidnight()
			duration := time.Until(nextMidnight)

			time.Sleep(duration)

			resetAllUserLimits()

			log.Printf("Daily limits reset at midnight: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		}
	}()
}

func resetAllUserLimits() {
	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	todayDate := getTodayDate()

	for _, user := range userStore.users {
		if !user.IsMaster {
			user.SearchCount = 0
			user.ResetDate = todayDate
		}
	}
}

func checkAndResetUserLimit(user *User) {
	todayDate := getTodayDate()

	if user.ResetDate != todayDate {
		user.SearchCount = 0
		user.ResetDate = todayDate
	}
}

func initializeMasterKey() {
	masterKey := generateMasterKey()

	userStore.mu.Lock()
	userStore.masterKey = masterKey

	masterUser := &User{
		GoogleID:    "MASTER",
		Email:       "master@system",
		APIKey:      masterKey,
		SearchCount: 0,
		ResetDate:   getTodayDate(),
		DailyLimit:  -1,
		IsMaster:    true,
	}
	userStore.users["MASTER"] = masterUser
	userStore.mu.Unlock()

	fmt.Println("SYNAPİC SEARCH")
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("MASTER API KEY: %s\n", masterKey)
	fmt.Println("   → Unlimited access with all permissions")
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("Server: http://localhost:3000")
	fmt.Println("-------------------------------------------------------------")
}

func handleGoogleAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Invalid method"})
		return
	}

	var req struct {
		GoogleID string `json:"google_id"`
		Email    string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Invalid request"})
		return
	}

	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	if user, exists := userStore.users[req.GoogleID]; exists {
		checkAndResetUserLimit(user)
		json.NewEncoder(w).Encode(AuthResponse{
			Success: true,
			APIKey:  user.APIKey,
			Message: "Login successful",
		})
		return
	}

	apiKey := generateAPIKey()
	newUser := &User{
		GoogleID:    req.GoogleID,
		Email:       req.Email,
		APIKey:      apiKey,
		SearchCount: 0,
		ResetDate:   getTodayDate(),
		DailyLimit:  100,
		IsMaster:    false,
	}

	userStore.users[req.GoogleID] = newUser

	json.NewEncoder(w).Encode(AuthResponse{
		Success: true,
		APIKey:  apiKey,
		Message: "Account created with 100 daily searches",
	})
}

func handleRegenerateKey(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "API key required"})
		return
	}

	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	var targetUser *User
	for _, user := range userStore.users {
		if user.APIKey == apiKey {
			targetUser = user
			break
		}
	}

	if targetUser == nil {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Invalid API key"})
		return
	}

	if targetUser.IsMaster {
		json.NewEncoder(w).Encode(AuthResponse{Success: false, Message: "Cannot regenerate master key"})
		return
	}

	newAPIKey := generateAPIKey()
	targetUser.APIKey = newAPIKey

	json.NewEncoder(w).Encode(AuthResponse{
		Success: true,
		APIKey:  newAPIKey,
		Message: "API key regenerated",
	})
}

func handleUserStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "API key required"})
		return
	}

	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	var targetUser *User
	for _, user := range userStore.users {
		if user.APIKey == apiKey {
			targetUser = user
			break
		}
	}

	if targetUser == nil {
		json.NewEncoder(w).Encode(map[string]interface{}{"error": "Invalid API key"})
		return
	}

	checkAndResetUserLimit(targetUser)

	remaining := targetUser.DailyLimit - targetUser.SearchCount
	if targetUser.IsMaster {
		remaining = -1
	} else if remaining < 0 {
		remaining = 0
	}

	json.NewEncoder(w).Encode(StatsResponse{
		SearchesToday: targetUser.SearchCount,
		DailyLimit:    targetUser.DailyLimit,
		Remaining:     remaining,
	})
}

func validateAPIKey(apiKey string) (*User, bool) {
	userStore.mu.Lock()
	defer userStore.mu.Unlock()

	if apiKey == userStore.masterKey {
		if masterUser, exists := userStore.users["MASTER"]; exists {
			return masterUser, true
		}
	}

	for _, user := range userStore.users {
		if user.APIKey == apiKey {
			if user.IsMaster {
				return user, true
			}

			checkAndResetUserLimit(user)

			if user.SearchCount >= user.DailyLimit {
				return user, false
			}

			user.SearchCount++
			return user, true
		}
	}

	return nil, false
}

func handlePublicSearch(w http.ResponseWriter, r *http.Request, reader *bluge.Reader) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Query parameter 'q' is required"})
		return
	}

	results := performSearch(reader, query)

	json.NewEncoder(w).Encode(SearchResponse{
		Results: results,
		Total:   len(results),
	})
}

func handleProtectedSearch(w http.ResponseWriter, r *http.Request, reader *bluge.Reader) {
	w.Header().Set("Content-Type", "application/json")

	apiKey := r.URL.Query().Get("apikey")
	if apiKey == "" {
		apiKey = r.Header.Get("X-API-Key")
	}

	if apiKey == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "API key required"})
		return
	}

	user, valid := validateAPIKey(apiKey)
	if !valid {
		if user != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Daily search limit exceeded"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid or expired API key"})
		}
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Query parameter 'q' is required"})
		return
	}

	searchType := r.URL.Query().Get("type")
	var results []SearchResult

	if searchType == "image" {
		indexPath := filepath.Join("../../İndexs/image-data")
		config := bluge.DefaultConfig(indexPath)
		typeReader, err := bluge.OpenReader(config)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open image index"})
			return
		}
		defer typeReader.Close()
		results = performSearch(typeReader, query)
	} else if searchType == "news" {
		indexPath := filepath.Join("../../İndexs/news-data")
		config := bluge.DefaultConfig(indexPath)
		typeReader, err := bluge.OpenReader(config)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open news index"})
			return
		}
		defer typeReader.Close()
		results = performSearch(typeReader, query)
	} else {
		results = performSearch(reader, query)
	}

	remaining := user.DailyLimit - user.SearchCount
	if user.IsMaster {
		remaining = -1
	}

	json.NewEncoder(w).Encode(SearchResponse{
		Results:    results,
		Total:      len(results),
		TotalUsed:  user.SearchCount,
		Remaining:  remaining,
		DailyLimit: user.DailyLimit,
	})
}

func handleIndexStats(w http.ResponseWriter, indexPath string) {
	w.Header().Set("Content-Type", "application/json")

	var totalSize int64
	err := filepath.Walk(indexPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to calculate index size"})
		return
	}

	sizeMB := float64(totalSize) / (1024 * 1024)
	sizeGB := sizeMB / 1024

	json.NewEncoder(w).Encode(IndexStatsResponse{
		IndexPath: indexPath,
		SizeBytes: totalSize,
		SizeMB:    sizeMB,
		SizeGB:    sizeGB,
	})
}

func levenshteinDistance(s1, s2 string) int {
	len1 := len(s1)
	len2 := len(s2)

	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}

	matrix := make([][]int, len1+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
		matrix[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			min := matrix[i-1][j] + 1
			if matrix[i][j-1]+1 < min {
				min = matrix[i][j-1] + 1
			}
			if matrix[i-1][j-1]+cost < min {
				min = matrix[i-1][j-1] + cost
			}

			matrix[i][j] = min
		}
	}

	return matrix[len1][len2]
}

func similarityScore(s1, s2 string) float64 {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	if s1 == s2 {
		return 1.0
	}

	s1NoSpace := strings.ReplaceAll(s1, " ", "")
	s2NoSpace := strings.ReplaceAll(s2, " ", "")

	if s1NoSpace == s2NoSpace {
		return 0.95
	}

	maxLen := len(s1)
	if len(s2) > maxLen {
		maxLen = len(s2)
	}

	if maxLen == 0 {
		return 0
	}

	distance := levenshteinDistance(s1, s2)
	similarity := 1.0 - float64(distance)/float64(maxLen)

	if strings.Contains(s1, s2) || strings.Contains(s2, s1) {
		similarity += 0.2
		if similarity > 1.0 {
			similarity = 1.0
		}
	}

	return similarity
}

func normalizeturkish(s string) string {
	replacer := strings.NewReplacer(
		"ı", "i", "İ", "i",
		"ğ", "g", "Ğ", "g",
		"ü", "u", "Ü", "u",
		"ş", "s", "Ş", "s",
		"ö", "o", "Ö", "o",
		"ç", "c", "Ç", "c",
	)
	return replacer.Replace(strings.ToLower(s))
}

func tokenSimilarity(query, text string) float64 {
	queryTokens := strings.Fields(strings.ToLower(query))
	textTokens := strings.Fields(strings.ToLower(text))

	if len(queryTokens) == 0 {
		return 0
	}

	matchCount := 0
	totalScore := 0.0

	for _, qToken := range queryTokens {
		bestMatch := 0.0
		for _, tToken := range textTokens {
			score := similarityScore(qToken, tToken)
			if score > bestMatch {
				bestMatch = score
			}
		}
		totalScore += bestMatch
		if bestMatch > 0.7 {
			matchCount++
		}
	}

	return totalScore / float64(len(queryTokens))
}

func performSearch(reader *bluge.Reader, queryStr string) []SearchResult {
	queryStr = strings.TrimSpace(strings.ToLower(queryStr))
	queryNormalized := normalizeturkish(queryStr)

	matchQuery := bluge.NewMatchQuery(queryStr).SetField("url")
	matchQuery.SetBoost(10.0)

	fuzzyQuery := bluge.NewFuzzyQuery(queryStr).SetField("url")
	fuzzyQuery.SetFuzziness(2)
	fuzzyQuery.SetBoost(8.0)

	titleQuery := bluge.NewMatchQuery(queryStr).SetField("title")
	titleQuery.SetBoost(9.0)

	titleFuzzyQuery := bluge.NewFuzzyQuery(queryStr).SetField("title")
	titleFuzzyQuery.SetFuzziness(2)
	titleFuzzyQuery.SetBoost(7.0)

	descQuery := bluge.NewMatchQuery(queryStr).SetField("description")
	descQuery.SetBoost(5.0)

	descFuzzyQuery := bluge.NewFuzzyQuery(queryStr).SetField("description")
	descFuzzyQuery.SetFuzziness(2)
	descFuzzyQuery.SetBoost(4.0)

	wildcardQuery := bluge.NewWildcardQuery("*" + queryStr + "*").SetField("url")
	wildcardQuery.SetBoost(6.0)

	wildcardTitleQuery := bluge.NewWildcardQuery("*" + queryStr + "*").SetField("title")
	wildcardTitleQuery.SetBoost(5.0)

	queryWords := strings.Fields(queryStr)

	boolQuery := bluge.NewBooleanQuery()
	boolQuery.AddShould(matchQuery)
	boolQuery.AddShould(fuzzyQuery)
	boolQuery.AddShould(titleQuery)
	boolQuery.AddShould(titleFuzzyQuery)
	boolQuery.AddShould(descQuery)
	boolQuery.AddShould(descFuzzyQuery)
	boolQuery.AddShould(wildcardQuery)
	boolQuery.AddShould(wildcardTitleQuery)

	for _, word := range queryWords {
		if len(word) > 2 {
			wordFuzzy := bluge.NewFuzzyQuery(word).SetField("url")
			wordFuzzy.SetFuzziness(1)
			wordFuzzy.SetBoost(3.0)
			boolQuery.AddShould(wordFuzzy)

			wordTitleFuzzy := bluge.NewFuzzyQuery(word).SetField("title")
			wordTitleFuzzy.SetFuzziness(1)
			wordTitleFuzzy.SetBoost(4.0)
			boolQuery.AddShould(wordTitleFuzzy)
		}
	}

	boolQuery.SetMinShould(1)

	request := bluge.NewTopNSearch(500, boolQuery).
		WithStandardAggregations()

	result, err := reader.Search(context.Background(), request)
	if err != nil {
		return []SearchResult{}
	}

	var searchResults []SearchResult
	match, err := result.Next()
	for err == nil && match != nil {
		doc := SearchResult{}

		err = match.VisitStoredFields(func(field string, value []byte) bool {
			switch field {
			case "url":
				doc.URL = string(value)
			case "title":
				doc.Title = string(value)
			case "description":
				doc.Description = string(value)
			}
			return true
		})

		doc.Score = calculateAdvancedRelevanceScore(doc.URL, doc.Title, doc.Description, queryStr, queryNormalized, match.Score)

		if doc.Score > 0.01 {
			searchResults = append(searchResults, doc)
		}

		match, err = result.Next()
	}

	if len(searchResults) == 0 {
		searchResults = fallbackSearch(reader, queryStr, queryNormalized)
	}

	sortByScore(searchResults)

	if len(searchResults) > 100 {
		searchResults = searchResults[:100]
	}

	return searchResults
}

func fallbackSearch(reader *bluge.Reader, query, queryNormalized string) []SearchResult {
	matchAllQuery := bluge.NewMatchAllQuery()
	request := bluge.NewTopNSearch(1000, matchAllQuery)

	result, err := reader.Search(context.Background(), request)
	if err != nil {
		return []SearchResult{}
	}

	var searchResults []SearchResult
	match, err := result.Next()

	for err == nil && match != nil {
		doc := SearchResult{}

		match.VisitStoredFields(func(field string, value []byte) bool {
			switch field {
			case "url":
				doc.URL = string(value)
			case "title":
				doc.Title = string(value)
			case "description":
				doc.Description = string(value)
			}
			return true
		})

		doc.Score = calculateAdvancedRelevanceScore(doc.URL, doc.Title, doc.Description, query, queryNormalized, 1.0)

		if doc.Score > 0.15 {
			searchResults = append(searchResults, doc)
		}

		match, err = result.Next()
	}

	sortByScore(searchResults)

	if len(searchResults) > 50 {
		searchResults = searchResults[:50]
	}

	return searchResults
}

func calculateAdvancedRelevanceScore(url, title, description, query, queryNormalized string, baseScore float64) float64 {
	score := baseScore
	urlLower := strings.ToLower(url)
	titleLower := strings.ToLower(title)
	descLower := strings.ToLower(description)
	queryLower := strings.ToLower(query)

	titleNormalized := normalizeturkish(titleLower)

	urlParts := strings.Split(strings.TrimPrefix(strings.TrimPrefix(urlLower, "http://"), "https://"), "/")
	domain := urlParts[0]
	domainNormalized := normalizeturkish(domain)

	if domain == queryLower || domain == queryLower+".com" || domain == "www."+queryLower+".com" {
		score *= 500.0
	}

	if domainNormalized == queryNormalized || domainNormalized == queryNormalized+".com" {
		score *= 400.0
	}

	if titleLower == queryLower || titleNormalized == queryNormalized {
		score *= 300.0
	}

	if strings.Contains(urlLower, queryLower) {
		score *= 100.0
	}

	if strings.HasPrefix(domain, queryLower) || strings.HasPrefix(domainNormalized, queryNormalized) {
		score *= 80.0
	}

	domainSimilarity := similarityScore(domain, queryLower)
	if domainSimilarity > 0.7 {
		score *= (50.0 * domainSimilarity)
	}

	titleSimilarity := similarityScore(titleLower, queryLower)
	if titleSimilarity > 0.6 {
		score *= (40.0 * titleSimilarity)
	}

	tokenSim := tokenSimilarity(queryLower, titleLower)
	if tokenSim > 0.6 {
		score *= (30.0 * tokenSim)
	}

	urlTokenSim := tokenSimilarity(queryLower, urlLower)
	if urlTokenSim > 0.6 {
		score *= (25.0 * urlTokenSim)
	}

	descTokenSim := tokenSimilarity(queryLower, descLower)
	if descTokenSim > 0.5 {
		score *= (10.0 * descTokenSim)
	}

	queryNoSpace := strings.ReplaceAll(queryLower, " ", "")
	titleNoSpace := strings.ReplaceAll(titleLower, " ", "")
	domainNoSpace := strings.ReplaceAll(domain, " ", "")

	if strings.Contains(domainNoSpace, queryNoSpace) {
		score *= 60.0
	}

	if strings.Contains(titleNoSpace, queryNoSpace) {
		score *= 50.0
	}

	if strings.Contains(titleLower, queryLower) {
		score *= 45.0
	}

	if strings.Contains(descLower, queryLower) {
		score *= 15.0
	}

	paramCount := strings.Count(urlLower, "?") + strings.Count(urlLower, "&") + strings.Count(urlLower, "=")
	if paramCount > 0 {
		score /= (1.0 + float64(paramCount)*0.3)
	}

	pathDepth := len(urlParts) - 1
	if pathDepth > 0 {
		score /= (1.0 + float64(pathDepth)*0.3)
	}

	if len(urlLower) > 100 {
		score /= 1.5
	}

	return score
}

func sortByScore(results []SearchResult) {
	for i := 0; i < len(results)-1; i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}

func main() {
	serverIP = getServerIP()
	if serverIP == "" {
		log.Println("Warning: Could not determine server IP - API endpoint will only accept localhost connections")
		serverIP = "127.0.0.1"
	}

	initializeMasterKey()
	startMidnightResetScheduler()

	indexPath := "../../İndexs/data"
	config := bluge.DefaultConfig(indexPath)
	reader, err := bluge.OpenReader(config)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	http.HandleFunc("/auth/google", corsMiddleware(handleGoogleAuth))
	http.HandleFunc("/auth/regenerate", corsMiddleware(handleRegenerateKey))
	http.HandleFunc("/user/stats", corsMiddleware(handleUserStats))

	http.HandleFunc("/api", apiIPRestriction(corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleSimpleSearch(w, r, reader)
	})))

	http.HandleFunc("/news", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleNewsSearch(w, r)
	}))

	http.HandleFunc("/images", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleImageSearch(w, r)
	}))

	http.HandleFunc("/ai", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleAIQuery(w, r, reader)
	}))

	http.HandleFunc("/api/search", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleProtectedSearch(w, r, reader)
	}))

	http.HandleFunc("/stats", corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleIndexStats(w, indexPath)
	}))

	fmt.Printf("\nAPI SECURITY: /api endpoint restricted to IP: %s\n", serverIP)
	log.Println("Server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleSimpleSearch(w http.ResponseWriter, r *http.Request, reader *bluge.Reader) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Query parameter 'q' is required"})
		return
	}

	searchType := r.URL.Query().Get("type")
	var results []SearchResult

	if searchType == "image" {
		indexPath := filepath.Join("../../İndexs/image-data")
		config := bluge.DefaultConfig(indexPath)
		typeReader, err := bluge.OpenReader(config)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open image index"})
			return
		}
		defer typeReader.Close()
		results = performSearch(typeReader, query)
	} else if searchType == "news" {
		indexPath := filepath.Join("../../İndexs/news-data")
		config := bluge.DefaultConfig(indexPath)
		typeReader, err := bluge.OpenReader(config)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open news index"})
			return
		}
		defer typeReader.Close()
		results = performSearch(typeReader, query)
	} else {
		results = performSearch(reader, query)
	}

	json.NewEncoder(w).Encode(SearchResponse{
		Results: results,
		Total:   len(results),
	})
}

func handleNewsSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Query parameter 'q' is required"})
		return
	}

	indexPath := filepath.Join("../../İndexs/news-data")
	config := bluge.DefaultConfig(indexPath)
	reader, err := bluge.OpenReader(config)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open news index"})
		return
	}
	defer reader.Close()

	results := performSearch(reader, query)

	json.NewEncoder(w).Encode(SearchResponse{
		Results: results,
		Total:   len(results),
	})
}

func handleImageSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(map[string]string{"error": "Query parameter 'q' is required"})
		return
	}

	indexPath := filepath.Join("../../İndexs/image-data")
	config := bluge.DefaultConfig(indexPath)
	reader, err := bluge.OpenReader(config)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to open image index"})
		return
	}
	defer reader.Close()

	results := performSearch(reader, query)

	json.NewEncoder(w).Encode(SearchResponse{
		Results: results,
		Total:   len(results),
	})
}
