// post_service.go
package services

import (
	"fmt"
	"my-go/models"
	"my-go/storage"

	"github.com/google/uuid"
)

func CreatePost(usern string, subr string, content string) error {
	post := &models.Post{
		ID:        uuid.New().String(),
		Author:    usern,
		Subreddit: subr,
		Upvote:    0,
		Downvote:  0,
		Content:   content,
		Comments:  []*models.Comment{},
	}
	return storage.AddPost(post)
}

func GetPosts(subr string) []*models.Post {
	var posts = []*models.Post{}

	for _, post := range storage.Storage.Posts {
		if post.Subreddit == subr {
			posts = append(posts, post)
		}
	}
	return posts
}

func ViewPosts(postid string) (*models.Post, error) {
	return storage.SeePost(postid)
}

func UpvotePost(postid string) error {
	if post, exists := storage.Storage.Posts[postid]; exists {
		post.Upvote++
		return nil
	}
	return fmt.Errorf("post not found")
}

func DownvotePost(postid string) error {
	if post, exists := storage.Storage.Posts[postid]; exists {
		post.Downvote++
		return nil
	}
	return fmt.Errorf("post not found")
}
