package models

import "sync"

type Comment struct {
	sync.Mutex

	ID       string
	Author   string
	ParentID string
	Content  string
	Upvote   int
	Downvote int
	Replies  []*Comment
}

// Subreddit string
