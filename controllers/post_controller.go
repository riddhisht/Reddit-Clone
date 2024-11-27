package controllers

import (
	"fmt"

	"my-go/services"
)

func HandleCreatePost(usern string, subr string, content string) {
	err := services.CreatePost(usern, subr, content)

	if err != nil {
		fmt.Println("Error Creating post ")
	} else {
		fmt.Println("Post created  ")
	}
}

func HandleGetFeed(subr string) []string {

	posts := services.GetPosts(subr)
	postIDs := make([]string, 0, len(posts))

	for _, post := range posts {
		postIDs = append(postIDs, post.ID)
		HandleGetPostMin(post.ID)
	}

	return postIDs
}
func HandleGetPostMin(postid string) {

	_, err := services.ViewPosts(postid)
	if err != nil {
		fmt.Println("Error getting post")
	}
	// fmt.Println("Karma Computed!! ", post.Upvote-post.Downvote)

	// else {
	// 	fmt.Println("_______________")
	// 	fmt.Println("Post ID: ", post.ID)
	// 	fmt.Println("Author: ", post.Author)
	// 	fmt.Println("Content: ", post.Content)
	// 	fmt.Println("Upvotes: ", post.Upvote, " Downvotes: ", post.Downvote)

	// }
}

func HandleGetPost(postID string) []string {
	post, err := services.ViewPosts(postID)
	if err != nil {
		fmt.Println("Error getting post:", err)
		return nil
	}
	fmt.Println("Karma Computed!! ", post.Upvote-post.Downvote)
	fmt.Println("Post ID:", post.ID)
	fmt.Println("Author:", post.Author)
	fmt.Println("Content:", post.Content)
	fmt.Println("Upvotes:", post.Upvote, "| Downvotes:", post.Downvote)
	fmt.Println("\nComments:")
	if len(post.Comments) > 0 {
		return HandleGetComments(post.ID, 1)
	} else {
		fmt.Println("No comments found.")
		return nil
	}
	// return nil
}

func HandleUpvotePost(postid string) {
	err := services.UpvotePost(postid)
	if err != nil {
		fmt.Println("Error Upvoting post ")
	} else {
		fmt.Println("Post upvoted ")
	}
}
func HandleDownVotePost(postid string) {
	err := services.DownvotePost(postid)
	if err != nil {
		fmt.Println("Error Downvoting post ")
	} else {
		fmt.Println("Post downvoted ")
	}
}
