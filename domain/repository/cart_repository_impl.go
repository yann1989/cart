package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yann1989/cart/domain/model"
)

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Cart{}).Error
}

func (c *CartRepository) FindCartByID(id int64) (*model.Cart, error) {
	cart := &model.Cart{}
	return cart, c.mysqlDb.First(cart, id).Error
}

func (c *CartRepository) FindAll(userID int64) ([]model.Cart, error) {
	ret := []model.Cart{}
	return ret, c.mysqlDb.Where("user_id = ?", userID).Find(&ret).Error
}

func (c *CartRepository) DeleteCartByID(id int64) error {
	return c.mysqlDb.Where("id = ?", id).Delete(&model.Cart{}).Error
}

func (c *CartRepository) CreateCart(cart *model.Cart) error {
	return c.mysqlDb.FirstOrCreate(cart, model.Cart{ProductID: cart.ProductID, SizeID: cart.SizeID, UserID: cart.UserID}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Model(cart).Update(cart).Error
}

func (c *CartRepository) CleanCart(userID int64) error {
	return c.mysqlDb.Where("user_id = ?", userID).Delete(&model.Cart{}).Error

}

func (c *CartRepository) IncrNum(cartID int64, num int64) error {
	return c.mysqlDb.Where("id = ?", cartID).UpdateColumn("num", gorm.Expr("num + ?", num)).Error
}

func (c *CartRepository) DecrNum(cartID int64, num int64) error {
	return c.mysqlDb.Where("id = ? AND num >= ?", cartID, num).UpdateColumn("num", gorm.Expr("num - ?", num)).Error
}

func NewCartRepository(mysqlDb *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: mysqlDb}
}
