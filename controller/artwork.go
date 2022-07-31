package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/saisesai/www.sxu.ink/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func ArtworkAddHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	rst, err := model.NewArtwork()
	if err != nil {
		L.WithError(err).Errorln("failed to create new artwork")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func ArtworkDeleteHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	err := model.DeleteArtwork(id)
	if err != nil {
		L.WithError(err).Errorln("failed to delete artwork")
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

func ArtworkGetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	rst, err := model.FindArtworkByUuid(id)
	if err != nil {
		L.WithError(err).Errorln("failed to get artwork")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rst)
}

func ArtworkSetHandler(ctx *gin.Context) {
	L := ctx.Value("L").(*logrus.Entry)
	id := ctx.Param("uuid")
	data := make(bson.M)
	err := ctx.BindJSON(&data)
	if err != nil {
		L.WithError(err).Errorln("failed to bind json")
		return
	}
	err = model.SetArtwork(id, data)
	if err != nil {
		L.WithError(err).Errorln("failed to set artwork")
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
