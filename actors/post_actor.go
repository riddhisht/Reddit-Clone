package actors

import (
	"my-go/controllers"

	"github.com/asynkron/protoactor-go/actor"
)

// Message Types
type CreatePost struct {
	Username  string
	Subreddit string
	Content   string
}
type UpvotePost struct {
	PostID string
}
type DownvotePost struct {
	PostID string
}
type GetFeed struct {
	Subreddit string
}
type GetFeedResponse struct {
	PostIDs []string
}

// PostActor
type PostActor struct{}

func (p *PostActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *CreatePost:
		controllers.HandleCreatePost(msg.Username, msg.Subreddit, msg.Content)

	case *UpvotePost:
		controllers.HandleUpvotePost(msg.PostID)

	case *DownvotePost:
		controllers.HandleDownVotePost(msg.PostID)

	case *GetFeed:
		feed := controllers.HandleGetFeed(msg.Subreddit)
		ctx.Respond(&GetFeedResponse{PostIDs: feed})
	}
}
