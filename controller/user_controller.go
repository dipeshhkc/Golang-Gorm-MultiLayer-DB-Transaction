package controller

import (
	"golang-transaction/model"
	"golang-transaction/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController : represent the user's controller contract
type UserController interface {
	AddUser(c *gin.Context)
	GetAllUser(c *gin.Context)
	TransferMoney(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

//NewUserController -> returns new user controller
func NewUserController(s service.UserService) UserController {
	return &userController{
		userService: s,
	}
}

func (u *userController) GetAllUser(c *gin.Context) {
	log.Print("[UserController]...get all Users")

	users, err := u.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *userController) AddUser(c *gin.Context) {
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

func (u *userController) TransferMoney(c *gin.Context) {
	log.Print("[UserController]...get all Users")

	var moneyTransfer model.MoneyTransfer
	if err := c.ShouldBindJSON(&moneyTransfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.userService.TransferMoney(moneyTransfer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while Transfering money"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Successfully Money Transferred"})
}
