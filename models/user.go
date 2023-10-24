package models

import (
	"gin_ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

// 查找用户名是否存在
func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

// 根据用户Id获取用户信息
func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

// 添加用户
func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

// 查询
//func GetUserTest(id int) (User, error) {
//	var user User
//	err := dao.Db.Where("id = ?", id).First(&user).Error
//	return user, err
//}
//
//// 查询列表
//func GetUserListTest() ([]User, error) {
//	var users []User
//	err := dao.Db.Where("id < ?", 4).Find(&users).Error
//	return users, err
//}
//
//// 添加
//func AddUser(username string) (int, error) {
//	user := User{Username: username}
//	err := dao.Db.Create(&user).Error
//	return user.Id, err
//}
//
//// 更新
//func UpdateUser(id int, username string) {
//	dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username)
//}
//
//// 删除
//func DeleteUser(id int) error {
//	err := dao.Db.Delete(&User{}).Error
//	return err
//}
