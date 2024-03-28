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
	InventoryRepo repository.InventoryRepository
}

func createInventory(inventoryRepo repository.InventoryRepository) *Inventory {
	return &Inventory{
		InventoryRepo: inventoryRepo,
	}
}

func (i *Inventory) AddItem(ctx context.Context, item models.Item, quantity int) (err error) {
	err = i.InventoryRepo.Add(ctx, item)
	return
}

func (i *Inventory) UpdateItem(ctx context.Context, item models.Item, quantity int) (err error) {
	err = i.InventoryRepo.Update(ctx, item)
	return
}

func (i *Inventory) RemoveItem(ctx context.Context, itemId int64) (err error) {
	err = i.InventoryRepo.Delete(ctx, itemId)
	return
}

func (i *Inventory) GetInventory(ctx context.Context) (items []models.Item, err error) {
	items, err = i.GetInventory(ctx)
	return
}
