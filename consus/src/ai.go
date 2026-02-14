package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/blugelabs/bluge"
)

type AIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AIRequest struct {
	Model       string      `json:"model"`
	Messages    []AIMessage `json:"messages"`
	Temperature float64     `json:"temperature"`
	MaxTokens   int         `json:"max_tokens"`
}

type AIChoice struct {
	Message AIMessage `json:"message"`
}

type AIResponse struct {
	Choices []AIChoice `json:"choices"`
}

type AIQueryResponse struct {
	Success        bool   `json:"success"`
	Answer         string `json:"answer"`
	Message        string `json:"message,omitempty"`
	RemainingCount int    `json:"remaining_count,omitempty"`
}

var MODELS = []string{
	"meta-llama/llama-3.1-405b-instruct:free",
	"nousresearch/hermes-3-llama-3.1-405b:free",
	"qwen/qwen3-235b-a22b:free",
	"openai/gpt-oss-120b:free",
	"meta-llama/llama-3.3-70b-instruct:free",
	"deepseek/deepseek-r1-0528:free",
	"google/gemini-2.0-flash-exp:free",
	"z-ai/glm-4.5-air:free",
	"google/gemma-3-27b-it:free",
	"nvidia/nemotron-3-nano-30b-a3b:free",
	"alibaba/tongyi-deepresearch-30b-a3b:free",
	"allenai/olmo-3.1-32b-think:free",
	"mistralai/mistral-small-3.1-24b-instruct:free",
	"cognitivecomputations/dolphin-mistral-24b-venice-edition:free",
	"openai/gpt-oss-20b:free",
	"mistralai/devstral-2512:free",
	"kwaipilot/kat-coder-pro:free",
	"qwen/qwen3-coder:free",
	"google/gemma-3-12b-it:free",
	"nvidia/nemotron-nano-12b-v2-vl:free",
	"nvidia/nemotron-nano-9b-v2:free",
	"nex-agi/deepseek-v3.1-nex-n1:free",
	"moonshotai/kimi-k2:free",
	"tngtech/deepseek-r1t2-chimera:free",
	"tngtech/deepseek-r1t-chimera:free",
	"tngtech/tng-r1t-chimera:free",
	"mistralai/mistral-7b-instruct:free",
	"qwen/qwen-2.5-vl-7b-instruct:free",
	"amazon/nova-2-lite-v1",
	"xiaomi/mimo-v2-flash:free",
	"google/gemma-3-4b-it:free",
	"google/gemma-3n-e4b-it:free",
	"qwen/qwen3-4b:free",
	"meta-llama/llama-3.2-3b-instruct:free",
	"google/gemma-3n-e2b-it:free",
	"arcee-ai/trinity-mini:free",
}

const OPENROUTER_API_URL = "https://openrouter.ai/api/v1/chat/completions"
const OPENROUTER_API_KEY = ""

func getOpenRouterAPIKey() string {
	if OPENROUTER_API_KEY != "" && OPENROUTER_API_KEY != "sk-or-v1-your-api-key-here" {
		return OPENROUTER_API_KEY
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		log.Println("Warning: OPENROUTER_API_KEY not set - AI features disabled")
		return ""
	}
	return apiKey
}

func callOpenRouterAPI(model string, messages []AIMessage) (string, error) {
	apiKey := getOpenRouterAPIKey()
	if apiKey == "" {
		return "", fmt.Errorf("OpenRouter API key not configured")
	}
	requestBody := AIRequest{
		Model:       model,
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   1000,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", OPENROUTER_API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HTTP-Referer", "https://localhost:3000")
	req.Header.Set("X-Title", "Synapic AI")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var aiResponse AIResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return "", err
	}

	if len(aiResponse.Choices) > 0 {
		return aiResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from AI")
}

func fetchAIResponse(query string, results []SearchResult, searchType string) (string, error) {
	resultsText := ""

	switch searchType {
	case "image":
		resultsText = "Image search results:\n\n"
		for i, result := range results {
			if i >= 10 {
				break
			}
			resultsText += fmt.Sprintf("%d. Title: %s\n   URL: %s\n   Description: %s\n\n",
				i+1, result.Title, result.URL, result.Description)
		}
	case "news":
		resultsText = "News search results:\n\n"
		for i, result := range results {
			if i >= 10 {
				break
			}
			resultsText += fmt.Sprintf("%d. Title: %s\n   URL: %s\n   Description: %s\n\n",
				i+1, result.Title, result.URL, result.Description)
		}
	default:
		resultsText = "Web search results:\n\n"
		for i, result := range results {
			if i >= 10 {
				break
			}
			resultsText += fmt.Sprintf("%d. Title: %s\n   URL: %s\n   Description: %s\n\n",
				i+1, result.Title, result.URL, result.Description)
		}
	}

	systemPrompt := `Siz Synapic AI, akıllı bir arama asistanısınız. Göreviniz şunlardır:
1. Sağlanan arama sonuçlarını analiz etmek
2. En alakalı bilgileri özetlemek
3. Arama sonuçlarına dayanarak kullanıcı sorularını yanıtlamak
4. Her zaman kullanıcının sorgusuyla aynı dilde yanıt vermek (Türkçe sorgu ise Türkçe yanıt verin)
5. Özlü ancak bilgilendirici olmak
6. Mümkün olduğunda kaynak başlıklarını belirterek bilgiye atıfta bulunmak
7. Markdown kullanmamak

Unutmayın: Siz Synapic AI'sınız, başka bir yapay zeka asistanı değilsiniz.`

	userPrompt := fmt.Sprintf("User Query: %s\n\n%s\n\nBased on these search results, please provide a helpful answer to the user's query. If the query is a question, answer it. If it's a topic, provide a summary of the key information.", query, resultsText)

	messages := []AIMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userPrompt},
	}

	for _, model := range MODELS {
		response, err := callOpenRouterAPI(model, messages)
		if err == nil && response != "" {
			return response, nil
		}
	}

	return "I'm sorry, all AI models are currently busy or unable to respond. Please try again shortly.", fmt.Errorf("all models failed")
}

func getIndexPathForType(searchType string) string {
	switch searchType {
	case "image":
		return filepath.Join("../../Datas/image-data")
	case "news":
		return filepath.Join("../../Datas/news-data")
	default:
		return filepath.Join("index")
	}
}

func handleAIQuery(w http.ResponseWriter, r *http.Request, reader *bluge.Reader) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("q")
	if query == "" {
		json.NewEncoder(w).Encode(AIQueryResponse{
			Success: false,
			Message: "Query parameter 'q' is required",
		})
		return
	}

	apiKey := r.URL.Query().Get("apikey")
	if apiKey == "" {
		json.NewEncoder(w).Encode(AIQueryResponse{
			Success: false,
			Message: "API key required for authentication",
		})
		return
	}

	user, valid := validateAPIKey(apiKey)
	if !valid {
		json.NewEncoder(w).Encode(AIQueryResponse{
			Success: false,
			Message: "Invalid API key or daily limit exceeded",
		})
		return
	}

	searchType := r.URL.Query().Get("type")
	if searchType != "web" && searchType != "image" && searchType != "news" && searchType != "" {
		searchType = "web"
	}
	if searchType == "" {
		searchType = "web"
	}

	var results []SearchResult

	if searchType == "web" {
		results = performSearch(reader, query)
	} else {
		indexPath := getIndexPathForType(searchType)
		config := bluge.DefaultConfig(indexPath)
		typeReader, err := bluge.OpenReader(config)
		if err != nil {
			remaining := user.DailyLimit - user.SearchCount
			if user.IsMaster {
				remaining = -1
			} else if remaining < 0 {
				remaining = 0
			}
			json.NewEncoder(w).Encode(AIQueryResponse{
				Success:        false,
				Message:        fmt.Sprintf("Failed to open %s index", searchType),
				RemainingCount: remaining,
			})
			return
		}
		defer typeReader.Close()
		results = performSearch(typeReader, query)
	}

	remaining := user.DailyLimit - user.SearchCount
	if user.IsMaster {
		remaining = -1
	} else if remaining < 0 {
		remaining = 0
	}

	if len(results) == 0 {
		json.NewEncoder(w).Encode(AIQueryResponse{
			Success:        false,
			Message:        "No search results found",
			RemainingCount: remaining,
		})
		return
	}

	aiAnswer, err := fetchAIResponse(query, results, searchType)
	if err != nil {
		json.NewEncoder(w).Encode(AIQueryResponse{
			Success:        false,
			Message:        "Failed to get AI response",
			Answer:         aiAnswer,
			RemainingCount: remaining,
		})
		return
	}

	json.NewEncoder(w).Encode(AIQueryResponse{
		Success:        true,
		Answer:         aiAnswer,
		RemainingCount: remaining,
	})
}
