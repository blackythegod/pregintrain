package webapi

import (
	"github.com/google/uuid"
	"time"
)

type Chat struct {
	Message     string    `json:"message"`
	MessageDate time.Time `json:"message_date"`
	Name        string    `json:"name"`
	Users       []*User   `json:"users"`
}
type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	CustomName   string    `json:"name"`
	Bio          string    `json:"bio"`
	PasswordHash string    `json:"password"`
	LastOnline   time.Time `json:"time_stamp"`
	Image        []byte    `json:"image"`
}
type Message struct {
	Id        uuid.UUID `json:"id"`
	ChatId    uuid.UUID `json:"chat_id"`
	UserId    uuid.UUID `json:"user_id"`
	Text      string    `json:"text"`
	TimeStamp time.Time `json:"time_stamp"`
}
