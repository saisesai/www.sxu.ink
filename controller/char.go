package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func CharAddHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	rst, err := model.NewChar()
	if err != nil {
		L.WithError(err).Errorln("failed to create new char")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func CharDeleteHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	err := model.DeleteChar(id)
	if err != nil {
		L.WithError(err).Errorln("failed to delete char")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"id":  id,
	})
}

func CharGetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	value := ctx.Param("value")
	rst, err := model.FindCharByValue(value)
	if err != nil {
		L.WithError(err).Errorln("failed to get char")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func CharSetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	data := make(bson.M)
	err := ctx.BindJSON(&data)
	if err != nil {
		L.WithError(err).Errorln("failed to bind json")
		return
	}
	err = model.SetChar(id, data)
	if err != nil {
		L.WithError(err).Errorln("failed to set char")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"id":  id,
	})
}

func CharSearchHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	query := make(bson.M)
	err := ctx.BindJSON(&query)
	if err != nil {
		L.WithError(err).Errorln("failed to bind json")
		return
	}
	rst, err := model.SearchChar(query)
	if err != nil {
		L.WithError(err).Errorln("failed to search char")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}
