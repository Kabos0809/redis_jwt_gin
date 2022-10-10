package Controller

import (
	"net/http"

	"github.com/kabos0809/redis_jwt_gin/Models"
	"github.com/gin-gonic/gin"
)

func (c Controller) SignUp(ctx *gin.Context) {
	var user Models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err":err})
		return
	} else {
		if err := c.Model.CreateUser(user.UName, user.Pass); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err":err})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{"status": "done!"})
		}
	}
}

func (c Controller) Login(ctx *gin.Context) {
	var user Models.User
	if err := ctx.Bind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	} else {
		if err := c.Model.CheckPassword(user.UName, user.Pass); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
	} else {
		tokens, err := Models.CreateJWTToken(user.UName, user.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err})
			return
		}
		ctx.JSON(http.StatusOK, tokens)
	}
	}
}

func (c Controller) Logout(ctx *gin.Context) {
	exp := ctx.MustGet("exp").(float64)
	token := ctx.MustGet("AccessToken").(string)
	err := Models.SetBlackList(token, int64(exp));
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err":err})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "Done"})
	}
}

func (c Controller) GetReToken(ctx *gin.Context) {
	ID := ctx.MustGet("ID").(float64)
	rt := ctx.MustGet("RefreshToken").(string)
	exp := ctx.MustGet("exp").(float64)
	token, err := c.Model.RefreshToken(uint64(ID), rt, int64(exp))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err":err})
		ctx.Abort()
	} else {
		ctx.JSON(http.StatusOK, token)
	}
}