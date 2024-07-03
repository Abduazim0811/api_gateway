package userapi

import (
	"Api/protos/userproto"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ClientUser userproto.UserServiceClient
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	var req userproto.UserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := u.ClientUser.CreateUser(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) Login(c *gin.Context) {
	var req userproto.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := u.ClientUser.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) GetByIDUser(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := u.ClientUser.GetByIDUser(ctx, &userproto.UserResponce{Id: int32(Id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (u *UserHandler) GetAllUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := u.ClientUser.GetAllUser(ctx, &userproto.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
