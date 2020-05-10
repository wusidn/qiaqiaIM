module github.com/wusidn/qiaqia/web

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.2
	github.com/golang/mock v1.4.3
	github.com/wusidn/qiaqia/dao v0.0.0
	github.com/wusidn/qiaqia/mockdao v0.0.0
)

replace (
	github.com/wusidn/qiaqia/dao => ../dao
	github.com/wusidn/qiaqia/mockdao => ../mockdao
)
