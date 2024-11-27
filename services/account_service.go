package services

import (
	"my-go/models"
	"my-go/storage"
)

func RegisterAccount(usern string, pword string) error {

	account := &models.Account{
		Username:   usern,
		Password:   pword,
		Subreddits: []string{},
	}
	return storage.AddAccount(usern, account)
}

func JoinSubreddit(usern string, subrname string) error {
	account, err := storage.GetAccount(usern)

	if err != nil {
		return err
	}
	subr, err := storage.GetSubreddit(subrname)

	if err != nil {
		return err
	}

	account.Subreddits = append(account.Subreddits, subrname)
	subr.Members = append(subr.Members, usern)

	return nil

}

func QuitSubreddit(usern string, subrname string) error {

	account, err := storage.GetAccount(usern)

	if err != nil {
		return err
	}
	subr, err := storage.GetSubreddit(subrname)

	if err != nil {
		return err
	}

	account.Subreddits = RemoveElement(account.Subreddits, subrname)
	subr.Members = RemoveElement(subr.Members, usern)

	return nil
}

func RemoveElement(slice []string, element string) []string {
	j := 0

	for i := 0; i < len(slice); i++ {
		if slice[i] != element {
			if i != j {
				slice[j] = slice[i]
			}
			j++
		}
	}
	return slice[:j]
}
