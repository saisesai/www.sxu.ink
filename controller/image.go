package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func ImageAddHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	rst, err := model.NewImage()
	if err != nil {
		L.WithError(err).Errorln("failed to create new image")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func ImageDeleteHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	err := model.DeleteImage(id)
	if err != nil {
		L.WithError(err).Errorln("failed to delete image")
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

func ImageGetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	rst, err := model.FindImageByUuid(id)
	if err != nil {
		L.WithError(err).Errorln("failed to get image")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func ImageSetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	data := make(bson.M)
	err := ctx.BindJSON(&data)
	if err != nil {
		L.WithError(err).Errorln("failed to bind json")
		return
	}
	err = model.SetImage(id, data)
	if err != nil {
		L.WithError(err).Errorln("failed to set image")
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
