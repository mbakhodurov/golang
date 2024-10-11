package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type TelegramService interface {
	SendMessage(message string) error
}

type telegramService struct {
	BotToken string
	ChatID   string
}

func NewTelegramService(botToken, ChatID string) TelegramService {
	return &telegramService{BotToken: botToken, ChatID: ChatID}
}

func (s *telegramService) SendMessage(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.BotToken)
	// Формируем параметры запроса
	data := url.Values{}
	data.Set("chat_id", s.ChatID)
	data.Set("text", message)

	// Отправляем POST запрос
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неудачный ответ от сервера: %s", resp.Status)
	}
	return nil
}
