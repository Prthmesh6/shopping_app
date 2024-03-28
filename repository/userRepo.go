package repository

import (
	"context"
	"sync"

	"github.com/Prthmesh6/shopping_app/models"
)

type UserRepository interface {
	Add(ctx context.Context, user models.User) (err error)
	Delete(ctx context.Context, userId int64) (err error)
	Update(ctx context.Context, user models.User) (err error)
	Getusers(ctx context.Context) (users []models.User, err error)
}

type userRepo struct {
	UsersMap map[int64]models.User
	Mutex    sync.RWMutex
}

func createUserRepo(UserRepo UserRepository) *userRepo {
	return &userRepo{
		UsersMap: make(map[int64]models.User),
		Mutex:    sync.RWMutex{},
	}
}

func (u *userRepo) Add(ctx context.Context, user models.User) (err error) {
	u.Mutex.Lock()
	u.UsersMap[user.Id] = user
	u.Mutex.Unlock()
	return
}

func (u *userRepo) Delete(ctx context.Context, userId int64) (err error) {
	u.Mutex.Lock()
	delete(u.UsersMap, userId)
	u.Mutex.Unlock()
	return
}

func (u *userRepo) Update(ctx context.Context, user models.User) (err error) {
	u.Mutex.Lock()
	u.UsersMap[user.Id] = user
	u.Mutex.Unlock()
	return
}

func (u *userRepo) Getusers(ctx context.Context) (users []models.User, err error) {
	u.Mutex.RLock()
	for _, val := range u.UsersMap {
		users = append(users, val)
	}
	u.Mutex.RUnlock()
	return
}
