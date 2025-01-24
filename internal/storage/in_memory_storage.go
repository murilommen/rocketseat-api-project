package storage

import (
	"errors"
	"github.com/google/uuid"
	"github.com/murilommen/rocketseat-api-project/internal/models"
)


type StorageInterface interface {
    Create(user models.User) (string, error)
	Get() ([]models.UserResponse, error)
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

func (us *UserStorage) Create(user models.User) (string, error) {
	id := uuid.New().String()

	if us.users == nil {
		return "", errors.New("database not initialized, could not create user")
	}

	if _, ok := us.users[id]; ok {
		return "", errors.New("user already exists in database")
	}

	us.users[id] = user

	return id, nil
}

func (us *UserStorage) Get() ([]models.UserResponse, error) {
	allUsers := make([]models.UserResponse, 0, len(us.users))

	for id, user := range us.users {
		allUsers = append(allUsers, models.UserResponse{Id: id, User: user})
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
	_, err := us.GetByID(id)
	if err != nil {
		return err
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