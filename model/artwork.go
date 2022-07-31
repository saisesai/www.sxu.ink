package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func NewArtwork() (bson.M, error) {
	var data = bson.M{
		"uuid":      uuid.New().String(),
		"create_at": time.Now(),
		"update_at": time.Now(),
		"delete_at": nil,
	}
	_, err := artworkCol.InsertOne(
		context.Background(),
		data,
	)
	if err != nil {
		return nil, err
	}
	logrus.Infof("artwork %v created!", data["uuid"])
	return data, nil
}

func DeleteArtwork(pUuid string) error {
	return artworkCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": bson.M{"delete_at": time.Now()}},
	)
}

func FindArtworkByUuid(pUuid string) (bson.M, error) {
	rst := make(bson.M)
	err := artworkCol.Find(
		context.Background(),
		bson.M{"uuid": pUuid},
	).One(rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func SetArtwork(pUuid string, pData bson.M) error {
	return artworkCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": pData},
	)
}
