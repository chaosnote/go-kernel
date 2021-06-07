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
var idx = 0
var store = map[string]map[int]Delegate{}

// Dispatch 觸發機制
//
// name : event name
//
// data :  any type
//
// go Dispatch('string', any)
func Dispatch(name string, data interface{}) {
	mu.Lock()
	if m, ok := store[name]; ok {
		for _, v := range m {
			v.On(data)
		}
	}
	mu.Unlock()
}

// Remove 移除指定事件
//
// name : event name
//
// id : identify id
func Remove(name string, id int) {
	mu.Lock()
	if m, ok := store[name]; ok {
		delete(m, id)
		store[name] = m
	}
	mu.Unlock()
}

// Register ...
// 字串
func Register(name string, d Delegate) int {
	var i = idx + 1
	mu.Lock()
	_, ok := store[name]
	if !ok {
		store[name] = map[int]Delegate{}
	}
	store[name][i] = d
	defer mu.Unlock()
	return i
}
