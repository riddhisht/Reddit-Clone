package actors

import (
	"my-go/controllers"

	"github.com/asynkron/protoactor-go/actor"
)

// Message Types
type CreateComment struct {
	PostID   string
	Username string
	Content  string
	ParentID string
}
type UpvoteComment struct {
	CommentID string
}
type DownvoteComment struct {
	CommentID string
}
type GetComments struct {
	ParentID string
	Indent   int
}
type GetCommentsResponse struct {
	CommentIDs []string
}

// CommentActor
type CommentActor struct{}

func (c *CommentActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *CreateComment:
		controllers.HandleCreateComment(msg.PostID, msg.Username, msg.Content, msg.ParentID)

	case *UpvoteComment:
		controllers.HandleUpvoteComment(msg.CommentID)

	case *DownvoteComment:
		controllers.HandleDownvoteComment(msg.CommentID)

	case *GetComments:
		comments := controllers.HandleGetComments(msg.ParentID, msg.Indent)
		ctx.Respond(&GetCommentsResponse{CommentIDs: comments})
	}
}
