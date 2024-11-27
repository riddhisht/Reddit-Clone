package services

import (
	"fmt"
	"my-go/models"
	"my-go/storage"

	"github.com/google/uuid"
)

// In services/comment_service.go
func CreateComment(postid string, usern string, content string, parentid string) (*models.Comment, error) {
	comment := &models.Comment{
		ID:       uuid.New().String(),
		Author:   usern,
		ParentID: parentid,
		Upvote:   0,
		Downvote: 0,
		Content:  content,
		Replies:  []*models.Comment{},
	}

	// Add comment to storage
	if err := storage.Storage.AddComment(comment); err != nil {
		return nil, err
	}

	// Handle parent relationship
	if parentComment, exists := storage.Storage.GetComment(parentid); exists {
		parentComment.Lock()
		defer parentComment.Unlock()
		parentComment.Replies = append(parentComment.Replies, comment)
	} else if post, err := storage.SeePost(postid); err == nil {
		post.Comments = append(post.Comments, comment)
	} else {
		return nil, fmt.Errorf("post not found")
	}

	return comment, nil
}

func UpvoteComment(commentid string) error {
	comment, exists := storage.Storage.GetComment(commentid)
	if !exists {
		return fmt.Errorf("comment not found")
	}

	comment.Lock()
	defer comment.Unlock()
	comment.Upvote++
	return nil
}
func DownvoteComment(commentid string) error {
	comment, exists := storage.Storage.GetComment(commentid)
	if !exists {
		return fmt.Errorf("comment not found")
	}

	comment.Lock()
	defer comment.Unlock()
	comment.Downvote++
	return nil
}

func GetComments(parentID string) []*models.Comment {
	var comments []*models.Comment
	for _, comment := range storage.Storage.Comments {
		if comment.ParentID == parentID {
			comments = append(comments, comment)
		}
	}
	return comments
}
