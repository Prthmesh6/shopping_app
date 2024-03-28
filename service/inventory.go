package service

import (
	"context"

	"github.com/Prthmesh6/shopping_app/models"
	"github.com/Prthmesh6/shopping_app/repository"
)

type InventoryManager interface {
	AddItem(ctx context.Context, item models.Item, quantity int) (err error)
	UpdateItem(ctx context.Context, item models.Item, quantity int) (err error)
	RemoveItem(ctx context.Context, itemId int64)
	GetInventory(ctx context.Context) (items []models.Item, err error)
}

type Inventory struct {
	InventoryRepo repository.InventoryRepo
}

func createInventory(inventoryRepo repository.InventoryRepo) *Inventory {
	return &Inventory{
		InventoryRepo: inventoryRepo,
	}
}

func (i *Inventory) AddItem(ctx context.Context, item models.Item, quantity int) (err error) {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) UpdateItem(ctx context.Context, item models.Item, quantity int) (err error) {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) RemoveItem(ctx context.Context, itemId int64) {
	panic("not implemented") // TODO: Implement
}

func (i *Inventory) GetInventory(ctx context.Context) (items []models.Item, err error) {
	panic("not implemented") // TODO: Implement
}
