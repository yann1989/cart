package handler

import (
	"cart/domain/model"
	"cart/domain/service"
	"cart/proto/cart"
	"context"
	"github.com/yann1989/common"
)

type Cart struct{
	CartService service.ICartService
}

func NewCart(cartService service.ICartService) *Cart {
	return &Cart{CartService: cartService}
}

func (c *Cart) AddCart(ctx context.Context, info *cart.CartInfo, resp *cart.ResponseAdd) error {
	tmp := &model.Cart{}
	err := common.SwapTo(info, tmp)
	if err != nil {
		return err
	}
	cartID, err := c.CartService.AddCart(tmp)
	if err != nil {
		return err
	}
	resp.CartId = cartID
	resp.Msg = "添加成功"
	return nil
}

func (c *Cart) CleanCart(ctx context.Context, clean *cart.Clean, response *cart.Response) error {
	err := c.CartService.CleanCart(clean.UserId)
	if err != nil {
		return err
	}
	response.Msg = "清空购物车成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, item *cart.Item, response *cart.Response) error {
	err := c.CartService.IncrNum(item.Id, item.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "操作成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, item *cart.Item, response *cart.Response) error {
	err := c.CartService.DecrNum(item.Id, item.ChangeNum)
	if err != nil {
		return err
	}
	response.Msg = "操作成功"
	return nil
}

func (c *Cart) DeleteItemByID(ctx context.Context, id *cart.CartID, response *cart.Response) error {
	err := c.CartService.DeleteCart(id.Id)
	if err != nil {
		return err
	}
	response.Msg = "操作成功"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, all *cart.CartFindAll, all2 *cart.CartAll) error {
	list , err := c.CartService.FindAllCart(all.UserId)
	if err != nil {
		return err
	}
	for _, info := range list {
		tmp := &cart.CartInfo{}
		if err = common.SwapTo(&info, tmp); err != nil {
			return err
		}
		all2.Infos = append(all2.Infos, tmp)
	}
	return nil
}


