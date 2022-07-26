package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt 三部分组成 header.payload.signature

// 定义过期时间
const TokenExpireDuration = time.Hour * 2

// 定义密钥
var MySecret = []byte("GXY")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt.StandardClaims 只包含了官方字段，我们也可以额外添加字段
type MyClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID uint64, username string) (token string, err error) {
	//创建一个我们自己声明的数据
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	fmt.Println(c)
	// 加密并获得玩真给的编码后的字符串
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)
	fmt.Println(token)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	//return token.SignedString(mySecret)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
