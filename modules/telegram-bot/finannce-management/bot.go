package finannce_management

import (
	"encoding/json"
	"fmt"
	"github.com/tnqbao/gau_assistant/config/gemini_api"
	"time"
)

func ParseMessageWithGemini(inputText string, chatID int64) (*Payment, error) {

	prompt := fmt.Sprintf(`
		Phân tích đoạn văn sau và trích xuất JSON:
		- "chat_id": %d
		- "category": danh mục (VD: "Thực phẩm", "Dịch vụ")
		- "object": mặt hàng (VD: "Bánh mì", "Tiền điện")
		- "price": số tiền (float)
		- "datePaid": ngày thanh toán (ISO 8601)

		Văn bản: "%s"
	`, chatID, inputText)
	aiClient := gemini_api.NewAIClient()
	resp, err := aiClient.GetResponse(prompt)
	if err != nil {
		return nil, err
	}

	var payment Payment
	err = json.Unmarshal([]byte(resp), &payment)
	if err != nil {
		return nil, fmt.Errorf("Lỗi xử lý JSON: %v", err)
	}

	if payment.DatePaid.IsZero() {
		payment.DatePaid = time.Now()
	}

	return &payment, nil
}
