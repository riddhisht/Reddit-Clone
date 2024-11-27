package controllers

import (
	"fmt"
	"my-go/services"
)

func HandleAccountRegistration(usern string, passw string) {
	err := services.RegisterAccount(usern, passw)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}

func HandleJoinSubreddit(usern string, subrname string) {
	err := services.JoinSubreddit(usern, subrname)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Subreddit Joined :) ")
	}

}

func HandleQuitSubreddit(usern string, subrname string) {
	err := services.QuitSubreddit(usern, subrname)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Subreddit Quitted :( ")
	}

}
