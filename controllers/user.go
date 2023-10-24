package controllers

import (
	"gin_ranking/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

// 注册
func (u UserController) Register(c *gin.Context) {
	// 接受用户名、密码、确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "密码和确认密码不相同")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "保存失败")
		return
	}
	ReturnSuccess(c, 200, "保存成功", "", 1)
}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// 登录
func (u UserController) Login(c *gin.Context) {
	// 接受用户名和密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 || user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	data := UserApi{Id: user.Id, Username: user.Username}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	ReturnSuccess(c, 0, "success", data, 1)

}

//func (u UserController) GetUserInfo(c *gin.Context) {
//	idstr := c.Param("id")
//	name := c.Param("name")
//
//	id, _ := strconv.Atoi(idstr)
//	user, _ := models.GetUserTest(id)
//
//	ReturnSuccess(c, 0, name, user, 1)
//}
//
//func (u UserController) AddUser(c *gin.Context) {
//	username := c.DefaultPostForm("username", "")
//	id, err := models.AddUser(username)
//	if err != nil {
//		ReturnError(c, 4002, "保存错误")
//		return
//	}
//	ReturnSuccess(c, 0, "保存成功", id, 1)
//}
//
//func (u UserController) UpdateUser(c *gin.Context) {
//	username := c.DefaultPostForm("username", "")
//	idstr := c.DefaultPostForm("id", "")
//	id, _ := strconv.Atoi(idstr)
//	models.UpdateUser(id, username)
//	ReturnSuccess(c, 0, "更新成功", true, 1)
//}
//
//func (u UserController) DeleteUser(c *gin.Context) {
//	idstr := c.DefaultPostForm("id", "")
//	id, _ := strconv.Atoi(idstr)
//	err := models.DeleteUser(id)
//	if err != nil {
//		ReturnError(c, 4002, "删除错误")
//		return
//	}
//	ReturnSuccess(c, 0, "删除成功", true, 1)
//}
//
//func (u UserController) GetList(c *gin.Context) {
//	//logger.Write("日志信息：", "user")
//	//defer func() {
//	//	if err := recover(); err != nil { //recover	回复程序
//	//		fmt.Println("捕获异常：", err)
//	//	}
//	//}()
//
//	num1 := 1
//	num2 := 0
//	num3 := num1 / num2
//	ReturnError(c, 4004, num3)
//}
//
//func (u UserController) GetUserListTest(c *gin.Context) {
//	users, err := models.GetUserListTest()
//	if err != nil {
//		ReturnError(c, 4004, "没有找到相关数据")
//		return
//	}
//	ReturnSuccess(c, 0, "获取数据成功", users, 1)
//}
