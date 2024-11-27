package main

import (
	"fmt"
	"math/rand"
	"time"

	"my-go/actors"

	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	fmt.Println("ProtoActor Reddit Simulation Started")

	// Initialize actor system
	system := actor.NewActorSystem()
	rootContext := system.Root

	// Create subreddit and account actors
	subredditProps := actor.PropsFromProducer(func() actor.Actor { return &actors.SubredditActor{} })
	accountProps := actor.PropsFromProducer(func() actor.Actor { return &actors.AccountActor{} })
	postProps := actor.PropsFromProducer(func() actor.Actor { return &actors.PostActor{} })
	commentProps := actor.PropsFromProducer(func() actor.Actor { return &actors.CommentActor{} })
	messageProps := actor.PropsFromProducer(func() actor.Actor { return &actors.MessageActor{} })

	subredditPids := []*actor.PID{}
	accountPids := []*actor.PID{}
	postPid := rootContext.Spawn(postProps)
	commentPid := rootContext.Spawn(commentProps)
	messagePid := rootContext.Spawn(messageProps)
	connectedAccounts := make(map[int]bool)

	// Create subreddits
	for i := 0; i < 10000; i++ {
		pid := rootContext.Spawn(subredditProps)
		subredditName := fmt.Sprintf("Subreddit-%d", i+1)
		rootContext.Send(pid, &actors.CreateSubreddit{
			Name:        subredditName,
			Description: fmt.Sprintf("Description of %s", subredditName),
		})
		subredditPids = append(subredditPids, pid)
	}

	// Register accounts
	for i := 0; i < 100000; i++ {
		pid := rootContext.Spawn(accountProps)
		username := fmt.Sprintf("User%d", i+1)
		rootContext.Send(pid, &actors.RegisterAccount{
			Username: username,
			Password: "password",
		})
		accountPids = append(accountPids, pid)
		connectedAccounts[i] = true
	}
	zipf := rand.NewZipf(rand.New(rand.NewSource(time.Now().UnixNano())), 1.2, 2.0, uint64(len(subredditPids)-1))
	total_posts_upvoted := 0
	total_posts_downvoted := 0
	total_comments_written := 0
	total_posts_created := 0
	total_reposts := 0
	total_comments_upvoted := 0
	total_comments_downvoted := 0
	total_subreddits_joined := 0
	total_subreddits_quitted := 0
	total_messages_sent := 0
	total_messages_replied := 0
	total_feed_checked := 0

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		subredditIndex := zipf.Uint64()
		subredditPid := subredditPids[subredditIndex]
		subredditName := fmt.Sprintf("Subreddit-%d", subredditIndex+1)
		uid := rand.Intn(len(accountPids)) + 1
		username := fmt.Sprintf("User%d", uid)
		action := rand.Intn(8)
		post_action := rand.Intn(7)
		if !connectedAccounts[uid] {
			continue
		}

		if i%10000 == 0 {
			for id := range connectedAccounts {
				connectedAccounts[id] = rand.Float32() < 0.75
			}
		}

		switch action {
		case 0: // Join subreddit
			total_subreddits_joined++
			rootContext.Send(subredditPid, &actors.JoinSubreddit{
				Username:  username,
				Subreddit: subredditName,
			})
		case 1: // Leave subreddit
			total_subreddits_quitted++
			rootContext.Send(subredditPid, &actors.LeaveSubreddit{
				Username:  username,
				Subreddit: subredditName,
			})
		case 2: // Create post
			content := fmt.Sprintf("Post content by %s", username)
			if rand.Float32() < 0.1 { // 10% chance of repost
				total_reposts++
				content = "Repost: " + content
			}
			total_posts_created++
			rootContext.Send(postPid, &actors.CreatePost{
				Username:  username,
				Subreddit: subredditName,
				Content:   content,
			})
		case 3:
			sender := fmt.Sprintf("User%d", rand.Intn(len(accountPids))+1)
			receiver := fmt.Sprintf("User%d", rand.Intn(len(accountPids))+1)
			if sender == receiver {
				// Skip if sender and receiver are the same
				continue
			}
			messageAction := rand.Intn(3)
			switch messageAction {
			case 0: // Send a message
				total_messages_sent++
				response, _ := rootContext.RequestFuture(messagePid, &actors.SendMessage{
					Sender:   sender,
					Receiver: receiver,
					Content:  fmt.Sprintf("Hello from %s to %s!", sender, receiver),
				}, 1*time.Second).Result()
				if _, ok := response.(*actors.SendMessageResponse); ok {
					fmt.Printf("Message sent from %s to %s with MessageID: %s\n", sender, receiver)
				}

			case 1: // Reply to a message
				total_messages_replied++
				response, _ := rootContext.RequestFuture(messagePid, &actors.SendMessage{
					Sender:   sender,
					Receiver: receiver,
					Content:  fmt.Sprintf("Replying from %s to %s!", sender, receiver),
				}, 1*time.Second).Result()
				if sendResponse, ok := response.(*actors.SendMessageResponse); ok {
					replyResponse, _ := rootContext.RequestFuture(messagePid, &actors.ReplyMessage{
						MessageID: sendResponse.MessageID,
						Content:   "This is a reply!",
					}, 1*time.Second).Result()
					if _, ok := replyResponse.(*actors.ReplyMessageResponse); ok {
						fmt.Printf("Reply sent")
					}
				}

			case 2: // Get direct messages
				total_feed_checked++
				rootContext.Send(messagePid, &actors.GetDirectMessages{
					Username: sender,
				})
				fmt.Printf("Fetched direct messages for user %s\n", sender)
			}

		default: // Post Masti
			response, _ := rootContext.RequestFuture(postPid, &actors.GetFeed{Subreddit: subredditName}, 1*time.Second).Result()
			if feedResponse, ok := response.(*actors.GetFeedResponse); ok {
				feed := feedResponse.PostIDs
				feed_len := len(feed)
				if feed_len == 0 {
					continue
				}

				switch post_action {
				case 0:
					// Upvote a post
					total_posts_upvoted++
					rootContext.Send(postPid, &actors.UpvotePost{
						PostID: feed[rand.Intn(feed_len)],
					})
				case 1:
					// Downvote a post
					total_posts_downvoted++
					rootContext.Send(postPid, &actors.DownvotePost{
						PostID: feed[rand.Intn(feed_len)],
					})
				case 2:
					// Interact with comments
					total_comments_written++
					rootContext.Send(commentPid, &actors.CreateComment{
						PostID:   feed[rand.Intn(feed_len)],
						Username: username,
						Content:  fmt.Sprintf("Comment-%d", rand.Intn(1000)),
						ParentID: feed[rand.Intn(feed_len)],
					})
				default:
					// Fetch comments
					commentResponse, _ := rootContext.RequestFuture(commentPid, &actors.GetComments{ParentID: feed[rand.Intn(feed_len)], Indent: 0}, 1*time.Second).Result()
					if commentResp, ok := commentResponse.(*actors.GetCommentsResponse); ok {
						// fmt.Println("Comments:", commentResp.CommentIDs)
						commies := commentResp.CommentIDs
						commie_len := len(commies)
						if commie_len == 0 {
							continue
						}
						fmt.Println("Insisde")
						toss := rand.Intn(100)
						if toss%3 == 0 {
							// Upvote a comment
							total_comments_upvoted++
							rootContext.Send(commentPid, &actors.UpvoteComment{
								CommentID: commentResp.CommentIDs[rand.Intn(commie_len)],
							})
						} else if toss%3 == 1 {
							// Downvote a comment
							total_comments_downvoted++
							rootContext.Send(commentPid, &actors.DownvoteComment{
								CommentID: commentResp.CommentIDs[rand.Intn(commie_len)],
							})
						} else {
							total_comments_written++
							rootContext.Send(commentPid, &actors.CreateComment{
								PostID:   commentResp.CommentIDs[rand.Intn(commie_len)],
								Username: username,
								Content:  fmt.Sprintf("Comment-%d", rand.Intn(1000)),
								ParentID: commentResp.CommentIDs[rand.Intn(commie_len)],
							})
						}

					}
				}

			}
		}
	}
	fmt.Println("Simulation Statistics:")
	fmt.Printf("Number of users: %d\n", 100000)
	fmt.Printf("Number of subredddits: %d\n", 10000)
	fmt.Printf("Total Posts Upvoted: %d\n", total_posts_upvoted)
	fmt.Printf("Total Posts Downvoted: %d\n", total_posts_downvoted)
	fmt.Printf("Total Comments Written: %d\n", total_comments_written)
	fmt.Printf("Total Posts Created: %d\n", total_posts_created)
	fmt.Printf("Total Reposts: %d\n", total_reposts)
	fmt.Printf("Total Comments Upvoted: %d\n", total_comments_upvoted)
	fmt.Printf("Total Comments Downvoted: %d\n", total_comments_downvoted)
	fmt.Printf("Total Subreddits Joined: %d\n", total_subreddits_joined)
	fmt.Printf("Total Subreddits Quitted: %d\n", total_subreddits_quitted)
	fmt.Printf("Total Messages Sent: %d\n", total_messages_sent)
	fmt.Printf("Total Messages Replied: %d\n", total_messages_replied)
	fmt.Printf("Total Feed Checked: %d\n", total_feed_checked)
}
