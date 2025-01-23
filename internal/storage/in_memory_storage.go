package storage

import (
	"errors"
	"github.com/google/uuid"
	"github.com/murilommen/rocketseat-api-project/internal/models"
)


type StorageInterface interface {
    Create(user models.User) error
	Get() ([]models.User, error)
    GetByID(id string) (models.User, error)
    Update(id string, user models.User) error
    Delete(id string) error
}

type UserStorage struct {
	users map[string]models.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users: make(map[string]models.User),
	}
}

func (us *UserStorage) Create(user models.User) error {
	id := uuid.New().String()
	us.users[id] = user
	return nil
}

func (us *UserStorage) Get() ([]models.User, error) {
	allUsers := make([]models.User, 0, len(us.users))

	for _, user := range us.users {
		allUsers = append(allUsers, user)
	}
	return allUsers, nil

}

func (us *UserStorage) GetByID(id string) (models.User, error) {
	user, exists := us.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}


func (us *UserStorage) Update(id string, user models.User) error {
	_, exists := us.users[id]
	if !exists {
		return errors.New("user not found")
	}
	us.users[id] = user
	return nil
}

func (us *UserStorage) Delete(id string) error {
	_, exists := us.users[id]
	if !exists {
		return errors.New("user not found")
	}
	delete(us.users, id)
	return nil
}