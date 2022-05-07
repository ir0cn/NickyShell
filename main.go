package main

import (
	"encoding/json"
	"fmt"
	"github.com/creack/pty"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"syscall"
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

type RouterFunc func(string, ...gin.HandlerFunc) gin.IRoutes
type HandleFunc func(*gin.Context)

type Router struct {
	Uri      string
	Method   RouterFunc
	Action   HandleFunc
	NeedAuth bool
}

// handler
func handler(f HandleFunc, needAuth bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if res, ok := r.(*Response); ok {
					ctx.JSON(400, res)
				} else {
					ctx.JSON(400, &Response{Code: 1000, Message: fmt.Sprintf("%v", r)})
				}
				ctx.Abort()
				return
			}
		}()

		if needAuth { // TODO: add auth check
			token := ctx.GetHeader("Authorization")
			fmt.Println("Token:", token)
			if _, ok := JwtVerify(token); !ok {
				ctx.JSON(401, &Response{Code: 1002, Message: "错误的认证信息"})
				ctx.Abort()
				return
			}
		}

		f(ctx)
	}
}

func main() {
	router := gin.Default()

	var routers = []Router{
		{"/api/v1/user/:userName/login", router.POST, UserLogin, false},
		{"/api/v1/assets/:assetId/connect", router.GET, WsConnect, false},
	}
	for _, r := range routers {
		r.Method(r.Uri, handler(r.Action, r.NeedAuth))
	}
	_ = router.Run(":8081")
}

var upgrade = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

type Message struct {
	Data string `json:"data"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

func WsConnect(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Sec-Websocket-Protocol")

	if _, ok := JwtVerify(token); !ok {
		ctx.JSON(401, &Response{Code: 1003, Message: "Invalid Token"})
		panic(nil)
	}

	ws, err := upgrade.Upgrade(ctx.Writer, ctx.Request, http.Header{
		"Sec-Websocket-Protocol": {token},
	})
	if err != nil {
		panic(&Response{Code: 10001, Message: fmt.Sprintf("upgrade.Upgrade failed: %v", err)})
	}
	cmd := exec.Command("bash")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		panic(&Response{Code: 10001, Message: fmt.Sprintf("pty.Start bash failed: %v", err)})
		return
	}

	defer func() { _ = ptmx.Close() }()
	defer func() { _ = ws.Close() }()

	go func() {
		defer func() { _ = ptmx.Close() }()
		defer func() { _ = ws.Close() }()
		for {
			_, data, err := ws.ReadMessage()
			if err != nil {
				fmt.Println("ws.read: ", err)
				break
			}
			msg := &Message{}
			if err := json.Unmarshal(data, msg); err != nil {
				break // not json data.
			}
			if msg.Data != "" {
				if _, err := ptmx.Write([]byte(msg.Data)); err != nil {
					fmt.Println("ptmx.write:", err)
					break
				}
			} else if msg.Cols != 0 {
				_ = pty.Setsize(ptmx, &pty.Winsize{Cols: uint16(msg.Cols), Rows: uint16(msg.Rows)})
			}
		}
		if err := cmd.Process.Signal(syscall.SIGHUP); err != nil {
			fmt.Println("signal err:", err)
		}
	}()

	for {
		data := make([]byte, 8192)
		if n, err := ptmx.Read(data); err != nil {
			fmt.Println("ptmx.read:", err)
			break
		} else if n > 0 {
			if err := ws.WriteMessage(websocket.BinaryMessage, data[:n]); err != nil {
				fmt.Println("ws.send error:", err)
				break
			}
		}
	}

	state, err := cmd.Process.Wait()
	if err != nil {
		fmt.Println("process.wait() error:", err)
	}
	if !state.Exited() {
		_ = cmd.Process.Kill()
	}
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func UserLogin(ctx *gin.Context) {
	defer func() { _ = ctx.Request.Body.Close() }()
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(&Response{Code: 1001, Message: "错误的提交数据"})
	}
	user := &LoginData{}
	if err := json.Unmarshal(data, user); err != nil {
		panic(&Response{Code: 1001, Message: "错误的提交数据"})
	}

	if user.Username != UserName || user.Password != Password {
		panic(&Response{Code: 1002, Message: "错误的用户名/口令"})
	}
	token := struct {
		Token string `json:"token"`
	}{}
	token.Token, err = JwtSign(user.Username)
	ctx.JSON(200, token)
}

func JwtSign(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(JWTKEY)
}

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
