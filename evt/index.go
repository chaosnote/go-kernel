package evt

import (
	"fmt"
	"sync"
)

/**
事件(觸發/註冊)
如事件數量過多，在思考是否修改為 chan 觸發監控
*/

// Delegate 註冊指定事件觸發後，由誰接手處理
type Delegate func(interface{})

var mu sync.Mutex
var idx = 0
var store = map[string]map[string]Delegate{}

// Dispatch 觸發機制
//
// name : event name
//
// data :  any type
//
// Dispatch('string', any)
func Dispatch(name string, data interface{}) {
	mu.Lock()
	if m, ok := store[name]; ok {
		for _, v := range m {
			go v(data) // race !?
		}
	}
	mu.Unlock()
}

// Remove 移除指定事件
//
// name : event name
//
// id : identify id
func Remove(name string, id string) {
	mu.Lock()
	if m, ok := store[name]; ok {
		delete(m, id)
		store[name] = m
	}
	mu.Unlock()
}

// Register ...
// 字串
func Register(name string, d Delegate) string {
	mu.Lock()
	idx = idx + 1
	id := fmt.Sprintf("%d", idx)
	_, ok := store[name]
	if !ok {
		store[name] = map[string]Delegate{}
	}
	store[name][id] = d
	defer mu.Unlock()
	return id
}
