package models

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
