package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type User struct {
	ID       string
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Product struct {
	ID          string
	ProductName string `json:"productName"`
	CategoryId  string `json:"categoryId"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Type        string `json:"type"`
	PhotoUrl    string `json:"photoUrl"`
}

type Category struct {
	ID           string
	CategoryName string `json:"categoryName"`
}

//
//type OrderHistory struct {
//	ID uint16
//	//OrderId   uint16    `json:"orderId"`
//	ProductId uint16    `json:"productId"`
//	UserId    uint16    `json:"userId"`
//	Quantity  uint16    `json:"quantity"`
//	DateTime  time.Time `json:"dateTime"`
//}
