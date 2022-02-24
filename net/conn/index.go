package conn

import (
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// Default ...
var Default = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	// 可調整為 設定檔 讀取
	Default.CheckOrigin = func(*http.Request) bool { return true }
}

//----------------------------------------------------------------------------------------------

/*
WebSocket ...

	conn.WebSocket(url.URL{
		Scheme: "ws",
		Host:   ":8080",
		Path:   "/",
	}).Dial()

*/
type WebSocket url.URL

// Dial ...
func (r WebSocket) Dial() (*websocket.Conn, error) {
	u := url.URL(r)
	d := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 5 * time.Second,
	}
	c, _, e := d.Dial(u.String(), nil)

	return c, e
}

// Accept ...
func Accept(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return Default.Upgrade(w, r, nil)
}
