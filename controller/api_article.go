package controller

import (
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateArticle(c *gin.Context)  {

	var articleReq models.ArticleCreateReq
	var article models.Article

	if c.ShouldBind(articleReq) == nil {
		article.Name = articleReq.Name
		article.Content = articleReq.Content

		_, err := models.CreateArticle(&article)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": tools.FAIL,
				"msg":    "稍后再试",
			})

		} else {

			c.JSON(http.StatusOK, gin.H{
				"status": tools.SUCCESS,
				"msg":    "文章创建成功",
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

func ArticleList(c *gin.Context)  {
	var articleListReq models.ArticleListReq
	if c.ShouldBind(articleListReq) == nil {


		result, err := models.ArticleList(&articleListReq)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": tools.SUCCESS,
				"msg":    "返回成功",
				"data":   result,
			})

		}else{
			c.JSON(http.StatusOK, gin.H{
				"status": tools.FAIL,
				"msg":    "请重试",
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