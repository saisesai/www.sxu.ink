package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func NewImage() (bson.M, error) {
	var data = bson.M{
		"uuid":      uuid.New().String(),
		"create_at": time.Now(),
		"update_at": time.Now(),
		"delete_at": nil,
	}
	_, err := imageCol.InsertOne(
		context.Background(),
		data,
	)
	if err != nil {
		return nil, err
	}
	logrus.Infof("image %v created!", data["uuid"])
	return data, nil
}

func DeleteImage(pUuid string) error {
	return imageCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": bson.M{"delete_at": time.Now()}},
	)
}

func FindImageByUuid(pUuid string) (bson.M, error) {
	rst := make(bson.M)
	err := imageCol.Find(
		context.Background(),
		bson.M{"uuid": pUuid},
	).One(rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func SetImage(pUuid string, pData bson.M) error {
	return imageCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": pData},
	)
}
