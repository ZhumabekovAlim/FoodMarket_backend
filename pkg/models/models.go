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
	ID       string `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Product struct {
	ID          string `json:"productId"`
	ProductName string `json:"productName"`
	CategoryId  string `json:"categoryId"`
	Price       string `json:"price"`
	Quantity    string `json:"quantity"`
	Type        string `json:"type"`
	PhotoUrl    string `json:"photoUrl"`
}

type Category struct {
	ID           string `json:"categoryId"`
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
