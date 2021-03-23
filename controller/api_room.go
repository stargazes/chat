package controller

import (
	"chat/middleware"
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//创建
func CreateRoom(c *gin.Context) {
	//接受post传参
	var room models.Room

	var req models.CreateReq

	if c.ShouldBind(&req) != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "输入参数有误",
		})
		return
	}

	claim := c.MustGet("claims").(middleware.CustomClaims) //获取中间件设置的token信息，主要拿userId
	userId, _ := strconv.Atoi(claim.ID)
	room.RoomHash = tools.GenerateRoomHash()
	room.UserId = int64(userId)
	room.Name = req.Name
	room.Desc = req.Desc

	_, err := models.FirstRoom("lyh")

	if err == nil {

		go func() {

			models.CreateRoom(&room) //创建MySQL记录

			tools.AddChatRoom(room.RoomHash, room.UserId)
		}()

		c.JSON(http.StatusOK, gin.H{
			"status": tools.SUCCESS,
			"msg":    "创建成功",
		})

	} else {

		//数据集合为空
		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "创建失败",
		})

	}

}

type AddForm struct {
	roomHash string `from:"roomHash" binding:"required"`
}

//加入
func AddRoom(c *gin.Context) {

	var addForm AddForm

	if c.ShouldBind(&addForm) != nil {
		//
		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "参数有误",
		})
		c.Abort()
	}

	claim := c.MustGet("claims").(middleware.CustomClaims) //获取中间件设置的token信息，主要拿userId
	userId, _ := strconv.Atoi(claim.ID)
	tools.AddChatRoom(addForm.roomHash, int64(userId))

	c.JSON(http.StatusOK, gin.H{
		"status": tools.SUCCESS,
		"msg":    "加入成功",
	})

}
