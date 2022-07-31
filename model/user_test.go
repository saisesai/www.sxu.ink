package model

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	un := &User{
		Username: "admin",
		Password: "admin",
		Nickname: "管理员",
		Group:    Admin,
	}
	err := NewUser(un)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindUserByUsername(t *testing.T) {
	uf, err := FindUserByUsername("admin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uf)
}

func TestUser_SetPassword(t *testing.T) {
	uf, err := FindUserByUsername("admin")
	if err != nil {
		t.Fatal(err)
	}
	err = uf.SetPassword("123456")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uf)
}

func TestUser_SetNickname(t *testing.T) {
	uf, err := FindUserByUsername("admin")
	if err != nil {
		t.Fatal(err)
	}
	err = uf.SetNickname("root")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uf)
}

func TestUser_SetGroup(t *testing.T) {
	uf, err := FindUserByUsername("admin")
	if err != nil {
		t.Fatal(err)
	}
	err = uf.SetGroup(Default)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(uf)
}
