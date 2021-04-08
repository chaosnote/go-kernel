package conn

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

// Upgrader ...
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	// 可調整為 設定檔 讀取
	Upgrader.CheckOrigin = func(*http.Request) bool { return true }
}

// WebSocket ...
// conn.WebSocket(url.URL{
// 	Scheme: "ws",
// 	Host:   ":8080",
// 	Path:   "/",
// }).Dial()
type WebSocket url.URL

// Dial ...
func (r WebSocket) Dial() (*websocket.Conn, error) {
	u := url.URL(r)
	c, _, e := websocket.DefaultDialer.Dial(u.String(), nil)

	if e != nil {
		return nil, e
	}

	return c, nil
}
