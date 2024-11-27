// message_service.go
package services

import (
	"fmt"
	"my-go/models"
	"my-go/storage"

	"github.com/google/uuid"
)

func SendMessage(sender string, receiver string, content string) (*models.Message, error) {
	message := &models.Message{
		ID:       uuid.New().String(),
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
	}
	storage.Storage.Messages[message.ID] = message
	return message, nil
}

func GetMessge(usern string) []*models.Message {
	var messages []*models.Message
	for _, mess := range storage.Storage.Messages {
		if mess.Receiver == usern {
			messages = append(messages, mess)
		}
	}
	return messages
}

func ReplyToMessage(messageID string, content string) (*models.Message, error) {
	if message, exists := storage.Storage.Messages[messageID]; exists {
		message.Replied = true
		reply := &models.Message{
			ID:       uuid.New().String(),
			Content:  content,
			Sender:   message.Receiver,
			Receiver: message.Sender,
			Replied:  true,
		}
		storage.Storage.Messages[reply.ID] = reply
		return reply, nil
	}
	return nil, fmt.Errorf("message not found")
}

func GetDirectMessages(receiver string) []*models.Message {
	var messages []*models.Message
	for _, message := range storage.Storage.Messages {
		if message.Receiver == receiver {
			messages = append(messages, message)
		}
	}
	return messages
}
