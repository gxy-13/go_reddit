package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwt 三部分组成 header.payload.signature

// 定义过期时间
const TokenExpireDuration = time.Hour * 2

// 定义密钥
var MySecret = []byte("GXY")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return MySecret, nil
}

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt.StandardClaims 只包含了官方字段，我们也可以额外添加字段
type MyClaims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 生成access token 和 refresh token aToken有效期短，rToken有效期长
func GenToken(userID uint64) (aToken, rToken string, err error) {
	//创建一个我们自己声明的数据
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "gxy",                                      // 签发人
		},
	}
	// 加密并获得玩真给的编码后的字符串,使用指定的secret签名并获得完整的编码后的字符串token
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)

	// rToken不需要存任何自定义数据
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * 30).Unix(),
		Issuer:    "gxy",
	}).SignedString(MySecret)

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

func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token 无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}

	// 从旧aToken中解析出claims数据
	var claims MyClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 当aToken是过期错误，并且rToken 没有过期时就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserID)
	}
	return
}
