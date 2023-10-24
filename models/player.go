package models

import (
	"gin_ranking/dao"
	"github.com/jinzhu/gorm"
)

type Player struct {
	Id          int    `json:"id"`          //选手Id
	Aid         int    `json:"aid"`         //所属排名活动Id
	Ref         string `json:"ref"`         //编号
	Nickname    string `json:"nickname"`    //昵称
	Declaration string `json:"declaration"` //宣言
	Avatar      string `json:"avatar"`      //头像
	Score       int    `json:"score"`       //分数
}

// 设置表名
func (Player) TableName() string {
	return "player"
}

// 查询参加活动的选手信息
func GetPlayers(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

// 根据选手Id获取选手信息
func GetPlayerInfo(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}

// 更新参赛选手分数字段，自增1
func UpdatePlayerScore(id int) {
	var player Player
	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
