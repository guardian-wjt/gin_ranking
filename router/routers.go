package router

import (
	"gin_ranking/config"
	"gin_ranking/controllers"
	"gin_ranking/pkg/logger"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//引用日志工具
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	// 路由组
	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)

		//user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		//user.POST("/list", controllers.UserController{}.GetList)
		//user.POST("/add", controllers.UserController{}.AddUser)
		//user.POST("/update", controllers.UserController{}.UpdateUser)
		//user.POST("/delete", controllers.UserController{}.DeleteUser)
		//user.POST("/list/test", controllers.UserController{}.GetUserListTest)
	}

	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}

	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}

	r.POST("/ranking", controllers.PlayerController{}.GetRanking)

	//order := r.Group("/order")
	//{
	//	order.POST("/list", controllers.OrderController{}.GetList)
	//}
	return r
}
