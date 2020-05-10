package web

import (
	"log"
	"errors"

	"github.com/gin-gonic/gin"
)

func (service *WebService)login(account, password string)(string, error){

	usersTable := service.db.GetTableUsers()

	userInfo, err := usersTable.QueryUserInfoByEmail(account)
	if err != nil{
		return "", err
	}

	if userInfo.EncodePassword != password{
		return "", errors.New("password mismatch")
	}

	token, err := generateToken(userInfo)

	if err != nil{
		log.Printf("generateToken fail [%v]", err.Error())
	}

	log.Printf("userInfo: %+v, token: %s", userInfo, token)

	return token, err
}


//登录检测中间件
func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){

		var token string = ""

		cookieToken, err := c.Request.Cookie("token")
		if err == nil{
			token = cookieToken.Value
		}

		if len(token) == 0 {
			headerToken := c.Request.Header["Token"]
			if len(headerToken) > 0{
				token = c.Request.Header["Token"][0]
			}
		}

		if len(token) == 0{
			token = c.Query("token")
		}

		if len(token) == 0{
			token = c.PostForm("token")
		}

		claims, err := parseToken(token)
		if err != nil{
			log.Printf("parse token fail [%v]", err.Error())
		}else{

			//将登录用户id保存在context
			c.Set("UserId", claims.UserId)
		}

		c.Next()
	}
}
