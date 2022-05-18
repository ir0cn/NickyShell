package service

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

var (
	JWTKEY   = []byte("HelloWorld")
	UserName = "admin"
	Password = "admin"
)

func init() {
	if key := os.Getenv("JWTKEY"); key != "" {
		JWTKEY = []byte(key)
	}
	if userName := os.Getenv("USERNAME"); userName != "" {
		UserName = userName
	}
	if passWord := os.Getenv("PASSWORD"); passWord != "" {
		Password = passWord
	}
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

// UserLogin 用户登录
func (svc *Service) UserLogin(ctx *gin.Context) {
	defer func() { _ = ctx.Request.Body.Close() }()
	data, err := ioutil.ReadAll(ctx.Request.Body)
	CheckError(err, "错误的提交数据")

	user := &LoginData{}
	err = json.Unmarshal(data, user)
	CheckError(err, "错误的提交数据")

	if user.Username != UserName || user.Password != Password {
		Throw("错误的用户名/口令")
	}

	token := &Token{}
	token.Token, err = JwtSign(user.Username)
	CheckError(err, "内部错误")
	ctx.JSON(200, token)
}

// Verify 验证是否是认证后的Token
func (svc *Service) Verify(ctx *gin.Context) bool {
	token := ctx.GetHeader("Authorization")
	fmt.Println("Token:", token)
	if _, ok := JwtVerify(token); !ok {
		ctx.JSON(401, "错误的认证信息")
		ctx.Abort()
		return false
	}
	return true
}

// JwtSign JWT 签名
func JwtSign(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(JWTKEY)
}

// JwtVerify JWT 验签名
func JwtVerify(tokenStr string) (string, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { return JWTKEY, nil })
	if err != nil {
		return "", false
	}
	if token.Valid {
		fmt.Println(token.Claims)
		claim := token.Claims.(jwt.MapClaims)
		u := claim["username"].(string)
		return u, true
	}
	return "", false
}
