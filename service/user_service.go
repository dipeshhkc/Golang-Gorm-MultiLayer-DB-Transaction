package service

import (
	"golang-transaction/model"
	"golang-transaction/repository"

	"gorm.io/gorm"
)

// UserService : represent the user's service contract
type UserService interface {
	Save(model.User) (model.User, error)
	GetAll() ([]model.User, error)
	WithTrx(*gorm.DB) userService
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(r repository.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

// WithTrx enables repository with transaction
func (u userService) WithTrx(trxHandle *gorm.DB) userService {
	u.userRepository = u.userRepository.WithTrx(trxHandle)
	return u
}

func (u userService) Save(user model.User) (model.User, error) {

	return u.userRepository.Save(user)
}

func (u userService) GetAll() ([]model.User, error) {

	return u.userRepository.GetAll()
}

func (u userService) IncrementMoney(receiver uint, amount float64) error {

	return u.userRepository.IncrementMoney(receiver, amount)
}

func (u userService) DecrementMoney(giver uint, amount float64) error {

	return u.userRepository.DecrementMoney(giver, amount)
}
