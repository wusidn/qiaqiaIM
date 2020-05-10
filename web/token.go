package web

import (
	"time"
	"errors"
	"github.com/wusidn/qiaqia/dao"
	"github.com/dgrijalva/jwt-go"
)

// errors
var (
    TokenExpired     error  = errors.New("Token is expired")
    TokenNotValidYet error  = errors.New("Token not active yet")
    TokenMalformed   error  = errors.New("That's not even a token")
    TokenInvalid     error  = errors.New("Couldn't handle this token:")
)

//token 体
type CustomClaims struct{
	UserId 		int		`json:"userId"`
	IsSystem	uint8	`json:"isSystem"`
	jwt.StandardClaims
}

var sigendKey []byte = []byte("wangqiaqia")

func generateToken(userInfo dao.UserInfo)(string, error){
	claims := CustomClaims{
		userInfo.Id,
		userInfo.IsSystem,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
            Issuer:    "wangbaichi",                    //签名的发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(sigendKey)

	return tokenString, err

}

func parseToken(tokenString string)(*CustomClaims, error){
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token)(interface{}, error){
		return sigendKey, nil
	})

	if err != nil{
		if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors & jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors & jwt.ValidationErrorExpired != 0 {
                return nil, TokenExpired
            } else if ve.Errors & jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}
