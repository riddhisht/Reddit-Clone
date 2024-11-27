package storage

import (
	"errors"
	"my-go/models"
)

func AddAccount(uname string, accs *models.Account) error {
	Storage.accountsMux.Lock()
	defer Storage.accountsMux.Unlock()

	if _, exists := Storage.Accounts[uname]; exists {
		return errors.New("account already exists")
	}
	Storage.Accounts[uname] = accs
	return nil
}

func GetAccount(uname string) (*models.Account, error) {
	Storage.accountsMux.Lock()
	defer Storage.accountsMux.Unlock()

	if account, exists := Storage.Accounts[uname]; exists {
		return account, nil
	}
	return nil, errors.New("account does not exist")
}

func AddSubreddit(name string, subrs *models.Subreddit) error {
	Storage.subredditsMux.Lock()
	defer Storage.subredditsMux.Unlock()

	if _, exists := Storage.Subreddits[name]; exists {
		return errors.New("subreddit already exists")
	}
	Storage.Subreddits[name] = subrs
	return nil
}

func GetSubreddit(name string) (*models.Subreddit, error) {
	Storage.subredditsMux.Lock()
	defer Storage.subredditsMux.Unlock()

	if subreddit, exists := Storage.Subreddits[name]; exists {
		return subreddit, nil
	}
	return nil, errors.New("subreddit does not exist")
}

func AddPost(post *models.Post) error {
	Storage.postsMux.Lock()
	defer Storage.postsMux.Unlock()

	if _, exists := Storage.Posts[post.ID]; exists {
		return errors.New("post already exists")
	}
	Storage.Posts[post.ID] = post
	return nil
}

func SeePost(postid string) (*models.Post, error) {
	Storage.postsMux.Lock()
	defer Storage.postsMux.Unlock()

	if post, exists := Storage.Posts[postid]; exists {
		return post, nil
	}
	return nil, errors.New("post does not exists")
}
