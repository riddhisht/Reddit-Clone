package models

type Post struct {
	ID        string
	Subreddit string
	Author    string
	Upvote    int
	Downvote  int
	Content   string
	Comments  []*Comment
}
