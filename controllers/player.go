package controllers

import (
	"gin_ranking/cache"
	"gin_ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type PlayerController struct {
}

// 获取活动选手信息
func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	rs, err := models.GetPlayers(aid, "id asc") // id 正序排序
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)

}

// 排行榜
func (p PlayerController) GetRanking(c *gin.Context) {
	//存储简单字符串
	//err := cache.Rdb.Set(cache.Rctx, "name", "zhangsan", 0).Err()
	//if err != nil {
	//	panic(err)
	//}

	// 先判断redis是否有我们需要的数据
	// 没有就从数据库中查询，找到并返回给前端、顺便保存到redis中
	// 下次获取数据的时候，先判断redis,有就直接从redis中取出数据，返给前端

	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	//Redis有序集合 优化 排行榜
	var redisKey string // key
	redisKey = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result() // id 与 分数
	if err == nil && len(rs) > 0 {
		var players []models.Player
		// 返给前端需要返回昵称，宣言等
		for _, value := range rs {
			//参赛选手的Id
			id, _ := strconv.Atoi(value)
			// 根据Id获取参赛选手的详情
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "success", players, 1)
		return
	}

	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		// 设置过期时间，防止Redis内存过大导致崩溃
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)

		ReturnSuccess(c, 0, "success", rsDb, 1)
		return
	}

	ReturnError(c, 4004, "没有获取到相关信息")
	return
}
