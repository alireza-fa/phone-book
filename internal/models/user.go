package models

type User struct {
	Id        uint64    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Contacts  []Contact `json:"contacts,omitempty"`
	CreatedAt string    `json:"created_at"`
}
