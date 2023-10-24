package controllers

import "github.com/gin-gonic/gin"

type OrderController struct {
}

//结构体获取Json数据
type Search struct {
	Name string `json:"name"`
	Cid  int    `json:"cid"`
}

func (o OrderController) GetList(c *gin.Context) {
	// POST方式获取请求参数
	//cid := c.PostForm("cid")
	//name := c.DefaultPostForm("name", "wangergou")

	// map获取Json数据
	//param := make(map[string]interface{})
	//err := c.BindJSON(&param)

	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		//ReturnSuccess(c, 0, param["name"], param["cid"], 1)

		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})

}
