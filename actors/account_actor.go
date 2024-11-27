package actors

import (
	"my-go/controllers"

	"github.com/asynkron/protoactor-go/actor"
)

type AccountActor struct{}

type RegisterAccount struct {
	Username string
	Password string
}

func (state *AccountActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *RegisterAccount:
		controllers.HandleAccountRegistration(msg.Username, msg.Password)
		break
		// default:
		// 	fmt.Println("Unknown message type in AccountActor")
	}
}
