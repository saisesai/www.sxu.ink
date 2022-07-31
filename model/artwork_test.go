package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewArtWork(t *testing.T) {
	rst, err := NewArtwork()
	if err != nil {
		panic(err)
	}
	fmt.Println(rst)
}

func TestDeleteArtwork(t *testing.T) {
	rst, err := NewArtwork()
	if err != nil {
		panic(err)
	}
	fmt.Println(rst)
	err = DeleteArtwork(rst["uuid"].(string))
	if err != nil {
		panic(err)
	}
}

func TestFindArtworkByUuid(t *testing.T) {
	id := "73fb57db-18d2-4b07-a61d-fbda75f96862"
	rst, err := FindArtworkByUuid(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(rst)
}

func TestSetArtwork(t *testing.T) {
	rst, err := NewArtwork()
	if err != nil {
		panic(err)
	}

	err = SetArtwork(rst["uuid"].(string), bson.M{
		"ti_ming":   "元简墓志",
		"zang_nian": "太和23年（499）3月8日",
	})
}

func TestSearchArtwork(t *testing.T) {
	rst, err := SearchArtwork(bson.M{"chu_tu_di": nil})
	if err != nil {
		panic(err)
	}
	for _, r := range rst {
		fmt.Println(r)
	}
}
