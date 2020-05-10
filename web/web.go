package web

import (
	"log"
	"sync"
	"net/http"
	"fmt"
	"github.com/wusidn/qiaqia/dao"
	"github.com/gin-gonic/gin"
)

type WebService struct {
	db dao.Dao
}

func New(db dao.Dao) *WebService{
	return &WebService{db}
}

func (service *WebService) Run(wg *sync.WaitGroup) {

	defer wg.Done()

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/", LoginMiddleware(), func(c *gin.Context) {

		userId := c.GetInt("UserId")

		userInfo, err := service.db.GetTableUsers().QueryUserInfo(userId)

		if err != nil{
			c.String(200, "未登录")
			return
		}

		c.String(200, fmt.Sprintf("hello %s", userInfo.Email ))
	})

	//login
	router.POST("/login", service.loginHandle)

	//regist
	router.POST("/regist", service.registHandle)

	router.Run(":8080")
}


type loginData struct{
	Token string	`json:"token"`
}

func (service *WebService)loginHandle(c *gin.Context){

	account := c.PostForm("account")
	password := c.PostForm("password")

	log.Printf("login: %s, %s", account, password)

	token, err := service.login(account, password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"message": "",
		})
		return
	}

	//设置cookie
	c.SetCookie("token", token, 60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"message": "登录成功",
		"data": loginData{
			Token: token,
		},
	})
}

func (service *WebService)registHandle(c *gin.Context){
	c.String(200, "regist")
}

