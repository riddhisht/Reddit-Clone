package services

import (
	"my-go/models"
	"my-go/storage"
)

func CreateSubreddit(name string, descrip string) error {
	subreddit := &models.Subreddit{
		Name:        name,
		Description: descrip,
		Members:     []string{},
	}
	return storage.AddSubreddit(name, subreddit)
}
