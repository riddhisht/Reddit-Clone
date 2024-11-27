package actors

import (
	"my-go/controllers"

	"github.com/asynkron/protoactor-go/actor"
)

type SubredditActor struct{}

type CreateSubreddit struct {
	Name        string
	Description string
}

type JoinSubreddit struct {
	Username  string
	Subreddit string
}

type LeaveSubreddit struct {
	Username  string
	Subreddit string
}

func (state *SubredditActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *CreateSubreddit:
		controllers.HandleSubredditCreation(msg.Name, msg.Description)
	case *JoinSubreddit:
		controllers.HandleJoinSubreddit(msg.Username, msg.Subreddit)
	case *LeaveSubreddit:
		controllers.HandleQuitSubreddit(msg.Username, msg.Subreddit)

		// default:
		// 	fmt.Println("Unknown message type in SubredditActor")
	}
}
