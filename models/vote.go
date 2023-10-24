package models

import (
	"gin_ranking/dao"
	"time"
)

type Vote struct {
	Id       int   `json:"id"`
	UserId   int   `json:"userId"`   //投票用户ID
	PlayerId int   `json:"playerId"` //选手ID
	AddTime  int64 `json:"addTime"`  //添加时间
}

func (Vote) TableName() string {
	return "vote"
}

// 用户是否已经投过票了
func GetVoteInfo(userId int, playerId int) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? and player_id = ?", userId, playerId).First(&vote).Error
	return vote, err
}

// 保存投票
func AddVote(userId int, playerId int) (int, error) {
	vote := Vote{UserId: userId, PlayerId: playerId, AddTime: time.Now().Unix()}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
