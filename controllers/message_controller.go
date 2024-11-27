package controllers

import (
	"fmt"
	"my-go/services"
)

func HandleSendMessage(sender string, receiver string, content string) string {
	mess, err := services.SendMessage(sender, receiver, content)
	if err != nil {
		fmt.Println("Error Sending Message ", err)
	} else {
		fmt.Println("Message Sent - ", mess)
	}
	return mess.ID
}
func HandleGetMessages(usern string) {
	messages := services.GetMessge(usern)

	for _, mess := range messages {
		fmt.Println(mess.Content)
	}
}
func HandleReplyMessage(messageid string, content string) string {
	reply, err := services.ReplyToMessage(messageid, content)

	if err != nil {
		fmt.Println("Error Replying Message ", err)
	} else {
		fmt.Println("Reply Sent - ", reply)
	}
	return reply.ID
}
func HandleGetDirectMessages(usern string) {
	messages := services.GetDirectMessages(usern)
	if len(messages) == 0 {
		fmt.Println("No direct messages found for user:", usern)
		return
	}

	fmt.Println("Direct Messages for", usern, ":")
	for _, mess := range messages {
		fmt.Printf("From: %s | Content: %s\n", mess.Sender, mess.Content)
	}
}
