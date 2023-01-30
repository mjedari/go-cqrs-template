package user

import (
	"context"
	"encoding/json"
	"github.com/mjedari/go-cqrs-template/app/providers/storage"
	"github.com/mjedari/go-cqrs-template/domain/user"
)

type IRepository interface {
	//
}
type UserRepository struct {
	//
	redisStorage storage.IStorage
}

func NewUserRepository(redisStorage storage.IStorage) *UserRepository {
	return &UserRepository{redisStorage: redisStorage}
}

func (r UserRepository) CreateUser(ctx context.Context, user *user.User) error {
	key, err := r.getKey(ctx, user)
	if err != nil {
		return err
	}

	userByte, err := json.Marshal(user)
	if err != nil {
		return err
	}

	if err = r.redisStorage.Insert(ctx, key, userByte); err != nil {
		return err
	}

	return nil
}

func (r UserRepository) GetUser(ctx context.Context, user *user.User) error {
	key := user.GetKey()

	userByte, err := r.redisStorage.Select(ctx, key)
	if err != nil {
		return err
	}

	err = json.Unmarshal(userByte, user)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepository) GetAllUsers(ctx context.Context) ([]user.User, error) {

	allUsersBytes, err := r.redisStorage.SelectAll(ctx, "user:*")
	if err != nil {
		return nil, err
	}

	var allUsers []user.User
	for _, ordersByte := range allUsersBytes {
		var user user.User
		if errM := json.Unmarshal(ordersByte, &user); errM != nil {
			return nil, errM
		}
		allUsers = append(allUsers, user)
	}

	return allUsers, nil
}

func (r UserRepository) UpdateUser(ctx context.Context, user *user.User) error {
	key := user.GetKey()

	userByte, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = r.redisStorage.Update(ctx, key, userByte); err != nil {
		return err
	}

	return nil
}

func (r UserRepository) getKey(ctx context.Context, user *user.User) (string, error) {
	err := r.addId(ctx, user)
	if err != nil {
		return "", err
	}
	key := user.GetKey()
	return key, nil
}

func (r UserRepository) addId(ctx context.Context, user *user.User) error {
	userId, err := r.redisStorage.GetNextId(ctx, user)
	if err != nil {
		return err
	}

	user.Id = uint(userId)
	return nil
}
