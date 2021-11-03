package conn

import (
	"net/http"
	"net/url"
	"sync"

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

var mMU sync.Mutex
var mStore = map[string]WebSocket{}

/*
Add2Store
加入暫存記錄

	key string
	item WebSocket

*/
func Add2Store(key string, item WebSocket) {
	mMU.Lock()
	defer mMU.Unlock()

	mStore[key] = item
}

/*
GetFromStore
取得暫存記錄

	key string

	return WebSocket,是否有值

*/
func GetFromStore(key string) (WebSocket, bool) {
	mMU.Lock()
	defer mMU.Unlock()

	_item, _ok := mStore[key]

	return _item, _ok
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
	c, _, e := websocket.DefaultDialer.Dial(u.String(), nil)

	return c, e
}

// Accept ...
func Accept(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	return Default.Upgrade(w, r, nil)
}
