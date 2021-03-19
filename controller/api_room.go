package controller

import (
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建
func CreateRoom(c *gin.Context) {
	//接受post传参
	var room models.Room
	//room.Name = c.PostForm("roomName")
	//room.Desc = c.PostForm("desc")
	room.Name = "lyh"
	room.Desc = "desc"
	room.RoomHash = tools.GenerateRoomHash()
	room.UserId = 1

	_, err := models.FirstRoom("lyh")

	if err != nil {

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
	roomHash string `from:"roomHash"`
	userId   int64  `from:"userId"`
}

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

	tools.AddChatRoom(addForm.roomHash,addForm.userId)
	
	c.JSON(http.StatusOK, gin.H{
		"status": tools.SUCCESS,
		"msg":    "加入成功",
	})

}
