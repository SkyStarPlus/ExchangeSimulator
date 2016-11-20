package Server

import (

)

type User interface {
	UserName() string
	UserId() string
	SetUserName(string)
}

type UserImpl struct {
	_name string
	_id string
}

func (user *UserImpl) UserName() string {
	return user._name
}

func (user *UserImpl) UserId() string {
	return user._id
}

//这是一个测试借口
func (user *UserImpl) SetUserName(name string) {
	user._name = name
}

func CreateUser(name string, id string) User {
	userImpl := new(UserImpl)
	userImpl._name = name
	userImpl._id = id
	return userImpl
}