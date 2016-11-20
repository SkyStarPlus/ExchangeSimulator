package Server

import (
)

type UserManager interface{
	AddUser(user User) (bool, error)
	FindUser(id string) (User, bool)
	Check(user User) bool
}

type UserManagerImpl struct {
	//TODO
	_userMap map[string]User
}

func (m *UserManagerImpl) AddUser(user User) (bool, error) {
	m._userMap[user.UserId()] = user
	return true, nil//TODO
}

func (m *UserManagerImpl) FindUser(id string) (User, bool) {
	user, exist := m._userMap[id]
	return user, exist
}

func (m *UserManagerImpl) Check(user User) bool{
	_, exist := m._userMap[user.UserId()]
	return exist//TODO
}

func (m *UserManagerImpl) Init() {
	//TODO
	//从数据库或者文件加载用户数据
	m._userMap = make(map[string]User)
}

func CreateUserManager() UserManager {
	userMgr := new(UserManagerImpl)
	userMgr.Init();
	return userMgr
}