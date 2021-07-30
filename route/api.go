package route

import (
	"chat/controller"
	"chat/middleware"
	"chat/models"
	"chat/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoutesForAll(e *gin.Engine)  {
	e.GET("/", func(c *gin.Context) {
		value := controller.TestRedis()
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
				"value": value,
			},
		)
	})

	e.GET("/init", InitModel) //初始化数据库
	e.POST("/user/register", controller.Register) //用户注册
	e.POST("/user/login", controller.Login)       //用户登录

	//需要登录才能访问的路由
	loginRoute := e.Group("/room")
	loginRoute.Use(middleware.JWTAuth())
	{

	}

	
}

/**
生成表格
*/
func InitModel(c *gin.Context)  {
	err :=tools.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Article{},&models.Category{},&models.Comment{})

	if err != nil {

		c.JSON(
			http.StatusOK,
			gin.H{
				"code":  http.StatusOK,
				"error": "数据库配置有误",
			},
		)
		return

	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": "数据库初始化成功",
		},
	)

	return
}
