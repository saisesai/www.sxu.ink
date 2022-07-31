package model

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/saisesai/www.sxu.ink/config"
	"github.com/sirupsen/logrus"
)

var db *qmgo.Database

var (
	userCol    *qmgo.Collection
	artworkCol *qmgo.Collection
	imageCol   *qmgo.Collection
	charCol    *qmgo.Collection
)

func init() {
	var err error
	client, err := qmgo.Open(context.Background(), &qmgo.Config{
		Uri:      *config.MongoUri,
		Database: *config.DbName,
	})
	if err != nil {
		logrus.WithError(err).Fatalln("failed to connect to db!")
	}
	db = client.Database
	userCol = db.Collection("users")
	artworkCol = db.Collection("artworks")
	imageCol = db.Collection("images")
	charCol = db.Collection("chars")
}
