// package storage

// import (
// 	"my-go/models"
// )

// var Accounts = make(map[string]*models.Account)
// var Subreddits = make(map[string]*models.Subreddit)
// var Posts = make(map[string]*models.Post)
// var Comments = make(map[string]*models.Comment)
// var Messages = make(map[string]*models.Message)
// In storage/storage.go
package storage

import (
	"my-go/models"
	"sync"
)

type SafeStorage struct {
	accountsMux   sync.RWMutex
	subredditsMux sync.RWMutex
	postsMux      sync.RWMutex
	commentsMux   sync.RWMutex

	Accounts   map[string]*models.Account
	Subreddits map[string]*models.Subreddit
	Posts      map[string]*models.Post
	Comments   map[string]*models.Comment
	Messages   map[string]*models.Message
}

var Storage = &SafeStorage{
	Accounts:   make(map[string]*models.Account),
	Subreddits: make(map[string]*models.Subreddit),
	Posts:      make(map[string]*models.Post),
	Comments:   make(map[string]*models.Comment),
	Messages:   make(map[string]*models.Message),
}

func (s *SafeStorage) AddComment(comment *models.Comment) error {
	s.commentsMux.Lock()
	defer s.commentsMux.Unlock()

	s.Comments[comment.ID] = comment
	return nil
}

func (s *SafeStorage) GetComment(commentID string) (*models.Comment, bool) {
	s.commentsMux.RLock()
	defer s.commentsMux.RUnlock()

	comment, exists := s.Comments[commentID]
	return comment, exists
}

// Update other methods similarly
