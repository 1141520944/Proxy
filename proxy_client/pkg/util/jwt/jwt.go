package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// token的过期时间
const TokenExpireDuration = time.Hour * 2

// token的sercet用于签名的字符串
var CustomSecret []byte = []byte("gin web")

type CustomClaims struct {
	jwt.RegisteredClaims        // 内嵌标准的声明
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
}

// GenToken 生成JWT
func GenToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		UserID:   userID,
		Username: username, // 自定义字段
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(TokenExpireDuration))
	claims.Issuer = "my-project"
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
