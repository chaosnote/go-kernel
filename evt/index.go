package evt

import (
	"sync"
)

/**
事件(觸發/註冊)
如事件數量過多，在思考是否修改為 chan 觸發監控
*/

// Delegate 註冊指定事件觸發後，由誰接手處理
type Delegate interface {
	On(interface{})
}

var mu sync.Mutex
var store = map[string]map[Delegate]bool{}

// Dispatch ...
// go Dispatch('string', any)
func Dispatch(name string, data interface{}) {
	mu.Lock()
	if m, ok := store[name]; ok {
		for d := range m {
			d.On(data)
		}
	}
	mu.Unlock()
}

// Remove ...
func Remove(name string, d Delegate) {
	mu.Lock()
	if m, ok := store[name]; ok {
		delete(m, d)
		store[name] = m
	}
	mu.Unlock()
}

// Register ...
// 字串
func Register(name string, d Delegate) {
	mu.Lock()
	_, ok := store[name]
	if !ok {
		store[name] = map[Delegate]bool{}
	}
	store[name][d] = true
	mu.Unlock()
}
