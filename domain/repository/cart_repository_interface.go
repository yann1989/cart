package repository

import "github.com/yann1989/cart/domain/model"

type ICartRepository interface {
	InitTable() error
	FindCartByID(id int64) (*model.Cart, error)
	FindAll(int64) ([]model.Cart, error)
	DeleteCartByID(id int64) error
	CreateCart(cart *model.Cart) error
	UpdateCart(cart *model.Cart) error

	CleanCart(int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}
