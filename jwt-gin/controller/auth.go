package controller

import (
	"gin/jwt-gin/model"
	"gin/jwt-gin/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAuth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var user UserAuth

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u := model.User{}
	u.Username = user.Username
	u.Password = user.Password

	_, err := u.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "registration success !",
	})
}

func Login(ctx *gin.Context) {
	var user UserAuth

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := model.User{}

	u.Username = user.Username
	u.Password = user.Password

	token, err := model.LoginCheck(u.Username, u.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "wrong password or username"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(ctx *gin.Context) {
	user_id, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := model.GetUserByID(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":     "success",
		"user_id": u,
	})
}
