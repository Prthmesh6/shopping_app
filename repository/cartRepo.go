package repository

import (
	"context"
	"sync"

	"github.com/Prthmesh6/shopping_app/models"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userId int64, item models.Item) (err error)
	// Delete(ctx context.Context, itemId int64) (err error)
	// Update(ctx context.Context, item models.Item) (err error)
	GetCart(ctx context.Context, userId int64) (items []models.Item, err error)
}

type cartRepo struct {
	CartMap map[int64]models.Cart
	Mutex   sync.RWMutex
}

func createCartRepo() *cartRepo {
	return &cartRepo{
		CartMap: make(map[int64]models.Cart),
		Mutex:   sync.RWMutex{},
	}
}

func (c *cartRepo) AddToCart(ctx context.Context, userId int64, item models.Item) (err error) {
	c.Mutex.Lock()
	c.CartMap[userId][item.ID] = item
	c.Mutex.Unlock()
	return
}

// Delete(ctx context.Context, itemId int64) (err error)
// Update(ctx context.Context, item models.Item) (err error)

func (c *cartRepo) GetCart(ctx context.Context, userId int64) (cart models.Cart, err error) {
	c.Mutex.RLock()
	cart = c.CartMap[userId]
	c.Mutex.RUnlock()
	return
}
