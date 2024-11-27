package models

type Message struct {
	ID       string
	Sender   string
	Receiver string
	Content  string
	Replied  bool
}

// Replied Bool
