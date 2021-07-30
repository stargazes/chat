package controller

import (
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(c *gin.Context)  {
	
	var categoryReq  models.CategoryCreateReq
	var category models.Category

	if c.ShouldBind(&categoryReq) == nil {
		category.CatName = categoryReq.CatName
		category.CatDesc = categoryReq.CatDesc

		_, err := models.CreateCategory(&category)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": tools.FAIL,
				"msg":    "稍后再试",
			})

		} else {

			c.JSON(http.StatusOK, gin.H{
				"status": tools.SUCCESS,
				"msg":    "分类创建成功",
			})

		}
		
	}else {
		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "输入参数有误",
		})
		return
	}
}
