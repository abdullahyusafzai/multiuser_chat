package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ChatRoom struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ChatMessage struct {
	ID         int       `json:"id"`
	ChatRoomID int       `json:"chat_room_id"`
	UserID     int       `json:"user_id"`
	Message    string    `json:"message"`
	created_at time.Time `json:"created_at"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./chatapp.db")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
