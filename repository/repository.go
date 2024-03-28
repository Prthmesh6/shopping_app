package repository

import (
	"context"

	"github.com/Prthmesh6/shopping_app/models"
)

type InventoryRepo interface {
	Add(ctx context.Context, item models.Item) (err error)
	Delete(ctx context.Context, itemId int64) (err error)
	Update(ctx context.Context, item models.Item) (err error)
	Getitems(ctx context.Context) (items []models.Item, err error)
}

type inventoryRepo struct {
	Items map[int64]models.Item
}

func (i *inventoryRepo) Add(ctx context.Context, item models.Item) (err error) {
	panic("not implemented") // TODO: Implement
}

func (i *inventoryRepo) Delete(ctx context.Context, itemId int64) (err error) {
	panic("not implemented") // TODO: Implement
}

func (i *inventoryRepo) Update(ctx context.Context, item models.Item) (err error) {
	panic("not implemented") // TODO: Implement
}

func (i *inventoryRepo) Getitems(ctx context.Context) (items []models.Item, err error) {
	panic("not implemented") // TODO: Implement
}
