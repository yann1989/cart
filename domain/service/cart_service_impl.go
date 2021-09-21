package service

import (
	"github.com/yann1989/cart/domain/model"
	"github.com/yann1989/cart/domain/repository"
)

type CartService struct {
	cartRepository repository.ICartRepository
}

func (c *CartService) AddCart(cart *model.Cart) (int64, error) {
	return cart.ID, c.cartRepository.CreateCart(cart)
}

func (c *CartService) DeleteCart(id int64) error {
	return c.cartRepository.DeleteCartByID(id)
}

func (c *CartService) UpdateCart(cart *model.Cart) error {
	return c.cartRepository.UpdateCart(cart)
}

func (c *CartService) FindCartByID(cartID int64) (*model.Cart, error) {
	return c.FindCartByID(cartID)
}

func (c *CartService) FindAllCart(userID int64) ([]model.Cart, error) {
	return c.cartRepository.FindAll(userID)
}

func (c *CartService) CleanCart(userID int64) error {
	return c.cartRepository.CleanCart(userID)
}

func (c *CartService) DecrNum(cartID int64, num int64) error {
	return c.cartRepository.DecrNum(cartID, num)

}

func (c *CartService) IncrNum(cartID int64, num int64) error {
	return c.cartRepository.IncrNum(cartID, num)

}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{cartRepository: cartRepository}
}
