package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID       uint16
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Product struct {
	ID          uint16
	ProductName string `json:"productName"`
	CategoryId  uint16 `json:"categoryId"`
	Price       uint16 `json:"price"`
	Quantity    uint16 `json:"quantity"`
	Type        string `json:"type"`
	PhotoUrl    string `json:"photoUrl"`
}

type Category struct {
	ID           int
	CategoryName string `json:"categoryName"`
}

type OrderHistory struct {
	ID uint16
	//OrderId   uint16    `json:"orderId"`
	ProductId uint16    `json:"productId"`
	UserId    uint16    `json:"userId"`
	Quantity  uint16    `json:"quantity"`
	DateTime  time.Time `json:"dateTime"`
}
