package model

import "time"

type Author struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Book struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	AuthorID uint   `json:"author_id"`
}

type BookPrice struct {
	BookID        uint    `json:"book_id"`
	RentalPrice   float64 `json:"rental_price"`
	PurchasePrice float64 `json:"purchase_price"`
}

type Transaction struct {
	ID              uint      `json:"id"`
	BookID          uint      `json:"book_id"`
	ClientID        uint      `json:"client_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	TakenAt         time.Time `json:"taken_at"`
	ReturnedAt      time.Time `json:"returned_at"`
}

type TransactionRequest struct {
	UserID    uint    `json:"user_id"`
	BookID    uint    `json:"book_id"`
	Price     float64 `json:"price"`
	Operation string  `json:"operation"` // может быть "rent" или "buy"
}
