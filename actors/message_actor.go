package actors

import (
	"my-go/controllers"

	"github.com/asynkron/protoactor-go/actor"
)

// Message Types
type SendMessage struct {
	Sender   string
	Receiver string
	Content  string
}
type SendMessageResponse struct {
	MessageID string
}
type ReplyMessage struct {
	MessageID string
	Content   string
}
type ReplyMessageResponse struct {
	ReplyID string
}
type GetMessages struct {
	Username string
}
type GetDirectMessages struct {
	Username string
}

// MessageActor
type MessageActor struct{}

func (m *MessageActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *SendMessage:
		messageID := controllers.HandleSendMessage(msg.Sender, msg.Receiver, msg.Content)
		ctx.Respond(&SendMessageResponse{MessageID: messageID})

	case *ReplyMessage:
		replyID := controllers.HandleReplyMessage(msg.MessageID, msg.Content)
		ctx.Respond(&ReplyMessageResponse{ReplyID: replyID})

	case *GetMessages:
		controllers.HandleGetMessages(msg.Username)

	case *GetDirectMessages:
		controllers.HandleGetDirectMessages(msg.Username)
	}
}
