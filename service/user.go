package service

import (
	"context"

	"github.com/Prthmesh6/shopping_app/models"
	"github.com/Prthmesh6/shopping_app/repository"
)

type UserManager interface {
	Add(ctx context.Context, user models.User) (err error)
	Delete(ctx context.Context, userId int64) (err error)
	Update(ctx context.Context, user models.User) (err error)
	GetCart(ctx context.Context, userId int64) (users models.Cart, err error)
	AddToCart(ctx context.Context, userId int64, item models.Item) (err error)
}

type User struct {
	userRepo repository.UserRepository
	cartRepo repository.CartRepository
}

func createUser(userRepo repository.UserRepository, cartRepo repository.CartRepository) *User {
	return &User{
		userRepo: userRepo,
		cartRepo: cartRepo,
	}
}

func (u *User) Add(ctx context.Context, user models.User) (err error) {
	err = u.userRepo.Add(ctx, user)
	return
}

func (u *User) Delete(ctx context.Context, userId int64) (err error) {
	err = u.userRepo.Delete(ctx, userId)
	return
}

func (u *User) Update(ctx context.Context, user models.User) (err error) {
	err = u.userRepo.Update(ctx, user)
	return
}

func (u *User) GetCart(ctx context.Context, userId int64) (cart models.Cart, err error) {
	cart, err = u.cartRepo.GetCart(ctx, userId)
	return
}

func (u *User) AddToCart(ctx context.Context, userId int64, item models.Item) (err error) {
	u.cartRepo.AddToCart(ctx, userId, item)
	return
}
