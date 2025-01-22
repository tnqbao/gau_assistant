package gemini_assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type AIClient struct {
	APIKey string
	APIURL string
}

func NewAIClient() *AIClient {
	return &AIClient{
		APIKey: os.Getenv("GEMINI_API_KEY"),
		APIURL: os.Getenv("GEMINI_API_URL"),
	}
}

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Candidate struct {
	Content Content `json:"content"`
}

type AIResponse struct {
	Candidates []Candidate `json:"candidates"`
}

func (client *AIClient) GetResponse(input string) (string, error) {
	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": input + "(trả lời tối đa 2000 kí tự bằng ngôn ngữ của câu hỏi)"},
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", client.APIURL+"?key="+client.APIKey, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result AIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		return result.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("unexpected response format: %+v", result)
}
