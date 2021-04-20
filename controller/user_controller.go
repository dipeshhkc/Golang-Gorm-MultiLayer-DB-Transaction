package controller

import (
	"golang-transaction/model"
	"golang-transaction/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController : represent the user's controller contract
type UserController interface {
	AddUser(*gin.Context)
	GetAllUser(*gin.Context)
	TransferMoney(*gorm.DB) gin.HandlerFunc
}

type userController struct {
	userService service.UserService
}

//NewUserController -> returns new user controller
func NewUserController(s service.UserService) UserController {
	return userController{
		userService: s,
	}
}

func (u userController) GetAllUser(c *gin.Context) {
	log.Print("[UserController]...get all Users")

	users, err := u.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u userController) AddUser(c *gin.Context) {
	log.Print("[UserController]...add User")
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userService.Save(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while saving user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (u userController) TransferMoney(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Print("[UserController]...get all Users")

		txHandle := db.Begin()

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		var moneyTransfer model.MoneyTransfer
		if err := c.ShouldBindJSON(&moneyTransfer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := u.userService.WithTrx(txHandle).IncrementMoney(moneyTransfer.Receiver, moneyTransfer.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while incrementing money"})
			txHandle.Rollback()
			return
		}

		if err := u.userService.WithTrx(txHandle).DecrementMoney(moneyTransfer.Giver, moneyTransfer.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decrementing money"})
			txHandle.Rollback()
			return
		}

		if err := txHandle.Commit().Error; err != nil {
			log.Print("trx commit error: ", err)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "Successfully Money Transferred"})
	}
}
