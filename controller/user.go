package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/middleware/auth"
	"github.com/saisesai/www.sxu.ink/model"
	"github.com/sirupsen/logrus"
)

func Hello(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	claims, err := auth.ParseMapClaimsJwt(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "hello, " + claims["username"].(string)})
}

func UserRegisterHandler(ctx *gin.Context) {
	var (
		user model.User
		err  error
		L    = ctx.Value("L").(*logrus.Entry)
	)
	if err = ctx.ShouldBindJSON(&user); err != nil {
		L.WithError(err).Errorln("failed to bind json")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err = model.NewUser(&user); err != nil {
		L.WithError(err).Errorln("failed to create new user")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}

func UserLoginHandler(ctx *gin.Context) {
	var (
		user model.User
		err  error
		L    = ctx.Value("L").(*logrus.Entry)
	)
	if err = ctx.ShouldBindJSON(&user); err != nil {
		L.WithError(err).Errorln("failed to bind json")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	tmpUser, err := model.FindUserByUsername(user.Username)
	if err != nil {
		L.WithError(err).Errorln("failed to find user")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if user.Password != tmpUser.Password {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": errors.New("password error"),
		})
		return
	}

	token, err := auth.BuildMapClaimsJwt(user.Username, user.Password)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"token": token,
	})
}

func UserSetPassword(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := auth.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}
	password := tmpParams["password"]
	err = user.SetPassword(password)
	if err != nil {
		L.WithError(err).Errorln("failed to set password")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func UserSetNickName(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := auth.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}
	nickname := tmpParams["nickname"]
	err = user.SetNickname(nickname)
	if err != nil {
		L.WithError(err).Errorln("failed to set nickname")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func UserSetGroup(ctx *gin.Context) {
	var (
		err       error
		L         = ctx.Value("L").(*logrus.Entry)
		tmpParams = make(map[string]string)
	)

	token := ctx.GetHeader("Authorization")
	claims, err := auth.ParseMapClaimsJwt(token)
	if err != nil {
		L.WithError(err).Errorln("failed to generate token")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = ctx.ShouldBind(&tmpParams)
	if err != nil {
		L.WithError(err).Errorln("failed to get params")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user := model.User{
		Username: claims["username"].(string),
	}
	group := tmpParams["group"]
	err = user.SetGroup(group)
	if err != nil {
		L.WithError(err).Errorln("failed to set group")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
