package controllers

import (
	"fmt"
	"my-go/models"
	"my-go/services"
	"sync"
)

var (
	Commies    = make(map[string]*models.Comment)
	commiesMux sync.Mutex // Mutex to protect the Posts map
)

func HandleCreateComment(postid string, usern string, content string, parentid string) {

	_, err := services.CreateComment(postid, usern, content, parentid)
	if err != nil {
		fmt.Println("Error Writing Comment ", err)
	} else {
		fmt.Println("Comment Written ")
	}
}
func HandleUpvoteComment(commentid string) {
	err := services.UpvoteComment(commentid)
	if err != nil {
		fmt.Println("Error Upvoting comment ")
	} else {
		fmt.Println("Comment upvoted ")
	}
}
func HandleDownvoteComment(commentid string) {
	err := services.DownvoteComment(commentid)
	if err != nil {
		fmt.Println("Error Downvoting comment ")
	} else {
		fmt.Println("Comment downvoted ")
	}
}

func HandleGetComments(parentID string, indent int) []string {
	commentIDs := []string{}
	comments := services.GetComments(parentID)

	for _, comment := range comments {
		comment.Lock() // Lock the comment
		HandleprintComment(comment, indent)
		commentIDs = append(commentIDs, comment.ID)

		// Safely iterate over Replies
		if len(comment.Replies) > 0 {
			replies := HandleGetComments(comment.ID, indent+1)
			commentIDs = append(commentIDs, replies...)
		}
		comment.Unlock() // Unlock the comment
	}

	return commentIDs
}

func HandleprintComment(comment *models.Comment, indent int) {

	padding := ""
	for i := 0; i < indent; i++ {
		padding += "    " // Indentation for nested replies
	}
	fmt.Println("_______________")
	fmt.Printf("%sComment ID: %s\n", padding, comment.ID)
	fmt.Printf("%sAuthor: %s\n", padding, comment.Author)
	fmt.Printf("%sContent: %s\n", padding, comment.Content)
	fmt.Printf("%sUpvotes: %d | Downvotes: %d\n\n", padding, comment.Upvote, comment.Downvote)
}

//	func HandleGetComments(parentID string, indent int) {
//		comments := services.GetComments(parentID) // Fetch comments for the given parentID
//		for _, comment := range comments {
//			HandleprintComment(comment, indent) // Print the comment
//			if len(comment.Replies) > 0 {
//				HandleGetComments(comment.ID, indent+1) // Recursively fetch and print replies
//			}
//		}
//	}
