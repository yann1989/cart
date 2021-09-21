package service

import "cart/domain/model"

type ICartService interface {
	AddCart(cart *model.Cart) (int64, error)
	DeleteCart(id int64) error
	UpdateCart(cart *model.Cart) error
	FindCartByID(cartID int64) (*model.Cart, error)
	FindAllCart(userID int64) ([]model.Cart, error)
	CleanCart(userID int64)error
	DecrNum(cartID int64, num int64) error
	IncrNum(cartID int64,num int64) error
}
