package model

import "time"

type Author struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Book struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	AuthorId uint   `json:"author_id"`
}

type BookHistory struct {
	Id         uint      `json:"id"`
	BookId     uint      `json:"book_id"`
	ClientId   uint      `json:"client_id"`
	TakenAt    time.Time `json:"taken_at"`
	ReturnedAt time.Time `json:"returned_at"`
}

type TakenBook struct {
	Id uint `json:"id"`
}

func (b *BookHistory) TableName() string {
	return "book_history"
}
