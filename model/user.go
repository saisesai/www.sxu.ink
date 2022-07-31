package model

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	Default = "default"
	Admin   = "admin"
)

type User struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Nickname string `json:"nickname" bson:"nickname"`
	Group    string `json:"group" bson:"group"`
}

func NewUser(pUser *User) error {
	rst, err := userCol.InsertOne(context.Background(), pUser)
	if err != nil {
		return err
	}
	logrus.Infof("user %v[%v] created!", pUser.Username, rst.InsertedID)
	return nil
}

func FindUserByUsername(pUsername string) (*User, error) {
	uf := &User{}
	err := userCol.Find(context.Background(), bson.M{"username": pUsername}).One(uf)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

func (u *User) SetPassword(pNewPassword string) error {
	err := userCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"password": pNewPassword}},
	)
	u.Password = pNewPassword
	return err
}

func (u *User) SetNickname(pNewNickname string) error {
	err := userCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"nickname": pNewNickname}},
	)
	u.Nickname = pNewNickname
	return err
}

func (u *User) SetGroup(pNewGroup string) error {
	err := userCol.UpdateOne(
		context.Background(),
		bson.M{"username": u.Username},
		bson.M{"$set": bson.M{"group": pNewGroup}},
	)
	u.Group = pNewGroup
	return err
}
