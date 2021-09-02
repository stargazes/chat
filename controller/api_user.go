package controller

import (
	"chat/models"
	"github.com/gin-gonic/gin"
)

type LoginResult struct {
	Token string `json:"token"`
	models.User
}

func Register(c *gin.Context) {


}

//用户登录
func Login(c *gin.Context) {

}

//生成token
func generateToken(c *gin.Context, user models.User) {
	//j := &middleware.JWT{
	//	[]byte("newtrekWang"),
	//}
	//claims := middleware.CustomClaims{
	//	string(user.ID),
	//	user.Username,
	//	user.Phone,
	//	jwt.StandardClaims{
	//		NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
	//		ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
	//		Issuer:    "newtrekWang",                   //签名的发行者
	//	},
	//}
	//
	//token, err := j.CreateToken(claims)
	//
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": -1,
	//		"msg":    err.Error(),
	//	})
	//	return
	//}
	//
	//log.Println(token)
	//
	//data := LoginResult{
	//	User:  user,
	//	Token: token,
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"status": tools.SUCCESS,
	//	"msg":    "登录成功！",
	//	"data":   data,
	//})
	//return
}
