package evt

import (
	"fmt"
	"sync"
)

/*
事件(觸發/註冊)(全域)
如事件數量過多，在思考是否修改為 chan 觸發監控
*/

// Delegate 註冊指定事件觸發後，由誰接手處理
type Delegate func(interface{})

var mPool = New()

/*
Dispatch 觸發機制
	name : event name
	data :  any type

	evt.Dispatch('string', any)
*/
func Dispatch(name string, data interface{}) {
	mPool.Dispatch(name, data)
}

/*
Remove 移除指定事件
	name : event name
	id : identify id

	evt.Remove('string', 'string')
*/
func Remove(name string, id string) {
	mPool.Remove(name, id)
}

/*
Register ...
	name : event name
	delegate : 委派對像

	evt.Register('string', Delegate)
*/
func Register(name string, d Delegate) string {
	return mPool.Register(name, d)
}

//-------------------------------------------------------------------------------------------------

/*
pool ...
	@see
		1、https://stackoverflow.com/a/37242475
		2、https://blog.golang.org/maps

*/
type pool struct {
	mu    *sync.Mutex
	idx   int
	store map[string]map[string]Delegate
}

/*
Dispatch 觸發機制
	name : event name
	data :  any type
*/
func (v pool) Dispatch(name string, data interface{}) {
	v.mu.Lock()
	if m, ok := v.store[name]; ok {
		for _, f := range m {
			go f(data) // race !?
		}
	}
	v.mu.Unlock()
}

/*
Remove 移除指定事件
	name : event name
	id : identify id
*/
func (v pool) Remove(name string, id string) {
	v.mu.Lock()
	if m, ok := v.store[name]; ok {
		delete(m, id)
		v.store[name] = m
	}
	v.mu.Unlock()
}

/*
Register ...
	name : event name
	delegate : 委派對像
*/
func (v *pool) Register(name string, d Delegate) string {
	v.mu.Lock()
	v.idx = v.idx + 1
	id := fmt.Sprintf("id.%d", v.idx)
	_, ok := v.store[name]
	if !ok {
		v.store[name] = map[string]Delegate{}
	}
	v.store[name][id] = d
	defer v.mu.Unlock()
	return id
}

/*
New ...
*/
func New() pool {
	p := pool{
		mu:    &sync.Mutex{},
		store: map[string]map[string]Delegate{},
	}

	return p
}
