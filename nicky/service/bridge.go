package service

import (
	"encoding/json"
	"fmt"
	"github.com/creack/pty"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"os/exec"
	"syscall"
)

var upgrade = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

const (
	CMD_DATA = iota
	CMD_RESIZE
	CMD_HEALTHCHECK
)

type Message struct {
	Cmd  int    `json:"cmd"`
	Data string `json:"data"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

func (svc *Service) WsConnect(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Sec-Websocket-Protocol")

	if _, ok := JwtVerify(token); !ok {
		ctx.JSON(401, "Invalid Token")
		panic(nil)
	}

	ws, err := upgrade.Upgrade(ctx.Writer, ctx.Request, http.Header{
		"Sec-Websocket-Protocol": {token},
	})
	if err != nil {
		panic(fmt.Sprintf("upgrade.Upgrade failed: %v", err))
	}
	cmd := exec.Command("bash")
	ptmx, err := pty.Start(cmd)
	if err != nil {
		panic(fmt.Sprintf("pty.Start bash failed: %v", err))
		return
	}

	defer func() { _ = ptmx.Close() }()
	defer func() { _ = ws.Close() }()

	go func() {
		defer func() { _ = ptmx.Close() }()
		defer func() { _ = ws.Close() }()

	ForLoop:
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
			switch msg.Cmd {
			case CMD_DATA:
				if _, err := ptmx.Write([]byte(msg.Data)); err != nil {
					fmt.Println("ptmx.write:", err)
					break ForLoop
				}
			case CMD_RESIZE:
				_ = pty.Setsize(ptmx, &pty.Winsize{Cols: uint16(msg.Cols), Rows: uint16(msg.Rows)})
			case CMD_HEALTHCHECK:
				// do nothing
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
