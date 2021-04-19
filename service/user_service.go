package service

import (
	"golang-transaction/model"
	"golang-transaction/repository"
)

// UserService : represent the user's service contract
type UserService interface {
	Save(model.User) (model.User, error)
	GetAll() ([]model.User, error)
	TransferMoney(model.MoneyTransfer) error
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (u *userService) Save(user model.User) (model.User, error) {

	return u.userRepository.Save(user)
}

func (u *userService) GetAll() ([]model.User, error) {

	return u.userRepository.GetAll()
}

func (u *userService) TransferMoney(MoneyTransfer model.MoneyTransfer) error {
	err := u.userRepository.IncrementMoney(MoneyTransfer.Receiver, MoneyTransfer.Amount)
	err = u.userRepository.DecrementMoney(MoneyTransfer.Giver, MoneyTransfer.Amount)
	return err
}
