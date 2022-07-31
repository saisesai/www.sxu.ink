package model

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func NewChar() (bson.M, error) {
	var data = bson.M{
		"uuid":      uuid.New().String(),
		"create_at": time.Now(),
		"update_at": time.Now(),
		"delete_at": nil,
	}
	_, err := charCol.InsertOne(
		context.Background(),
		data,
	)
	if err != nil {
		return nil, err
	}
	logrus.Infof("char %v created!", data["uuid"])
	return data, nil
}

func DeleteChar(pUuid string) error {
	return charCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": bson.M{"delete_at": time.Now()}},
	)
}

func FindCharByUuid(pUuid string) (bson.M, error) {
	rst := make(bson.M)
	err := charCol.Find(
		context.Background(),
		bson.M{"uuid": pUuid},
	).One(rst)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func SetChar(pUuid string, pData bson.M) error {
	return charCol.UpdateOne(
		context.Background(),
		bson.M{"uuid": pUuid},
		bson.M{"$set": pData},
	)
}

func SearchChar(pQuery bson.M) ([]bson.M, error) {
	var batch []bson.M
	err := charCol.Find(context.Background(), pQuery).All(&batch)
	if err != nil {
		return nil, err
	}
	return batch, nil
}
