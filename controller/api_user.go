package controller

import (
	"chat/middleware"
	"chat/models"
	"chat/tools"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type LoginResult struct {
	Token string `json:"token"`
	models.User
}

func Register(c *gin.Context) {

	var register models.RegisterReq
	var user models.User

	fmt.Println(c.PostForm("mobile"), c.PostForm("pwd"), c.PostForm("user_name"))
	if c.ShouldBind(&register) == nil {

		user.Phone = register.Phone
		user.Password = register.Pwd
		user.Username = register.UserName

		_, err := models.Insert(user)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": tools.FAIL,
				"msg":    "稍后再试",
			})

		} else {

			generateToken(c, user)

			return

		}

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "输入参数有误",
		})
		return
	}

}

//用户登录
func Login(c *gin.Context) {

	var loginReq models.LoginReq

	if c.ShouldBind(&loginReq) == nil {
		//绑定正常
		isPass, user, _, msg := models.LoginCheck(loginReq)

		if isPass {
			generateToken(c, user)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    msg,
		})


	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": tools.FAIL,
			"msg":    "输入参数有误",
		})
		return

	}

}

//生成token
func generateToken(c *gin.Context, user models.User) {
	j := &middleware.JWT{
		[]byte("newtrekWang"),
	}
	claims := middleware.CustomClaims{
		string(user.ID),
		user.Username,
		user.Phone,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": tools.SUCCESS,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}
