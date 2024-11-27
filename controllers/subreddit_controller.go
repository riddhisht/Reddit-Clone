package controllers

import (
	"fmt"
	"my-go/services"
)

func HandleSubredditCreation(name string, descrip string) {
	err := services.CreateSubreddit(name, descrip)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
