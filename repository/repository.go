package repository

import (
	"context"
	"sync"

	"github.com/Prthmesh6/shopping_app/models"
)

type InventoryRepository interface {
	Add(ctx context.Context, item models.Item) (err error)
	Delete(ctx context.Context, itemId int64) (err error)
	Update(ctx context.Context, item models.Item) (err error)
	Getitems(ctx context.Context) (items []models.Item, err error)
}

type inventoryRepo struct {
	Items map[int64]models.Item
	Mutex sync.RWMutex
}

func createInventory() *inventoryRepo {
	return &inventoryRepo{
		Items: make(map[int64]models.Item),
		Mutex: sync.RWMutex{},
	}
}

func (i *inventoryRepo) Add(ctx context.Context, item models.Item) (err error) {
	i.Mutex.Lock()
	i.Items[item.ID] = item
	i.Mutex.Unlock()
	return
}

func (i *inventoryRepo) Delete(ctx context.Context, itemId int64) (err error) {
	i.Mutex.Lock()
	delete(i.Items, itemId)
	i.Mutex.Unlock()
	return
}

func (i *inventoryRepo) Update(ctx context.Context, item models.Item) (err error) {
	i.Mutex.Lock()
	i.Items[item.ID] = item
	i.Mutex.Unlock()
	return
}

func (i *inventoryRepo) Getitems(ctx context.Context) (items []models.Item, err error) {
	i.Mutex.RLock()
	for _, val := range i.Items {
		items = append(items, val)
	}
	i.Mutex.RUnlock()
	return
}
