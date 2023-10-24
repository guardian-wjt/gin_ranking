package controllers

import (
	"gin_ranking/cache"
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct {
}

// 投票
func (v VoteController) AddVote(c *gin.Context) {
	//获取用户Id(投票人Id)	选手Id
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}

	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛选手不存在")
		return
	}

	// 是否已经投过票了
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "已投票了")
		return
	}

	// 上述检验通过后，就保存投票
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		//要更新参赛选手分数字段，自增1
		models.UpdatePlayerScore(playerId)

		// 同时要更新redis	- 确保排名时效一致
		var redisKey string // key
		redisKey = "ranking:" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))
		ReturnSuccess(c, 0, "投票成功了", rs, 1)
		return
	}
	ReturnError(c, 4004, "请联系管理员")
	return

}
