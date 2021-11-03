package evt

import (
	"fmt"
	"sync"
)

//-------------------------------------------------------------------------------------------------

type Callback func(interface{})

//-------------------------------------------------------------------------------------------------

/*
事件池
*/
type pool struct {
	mMu    sync.Mutex
	mStore map[string]Callback
}

/*
Dispatch
事件觸發

	key string
	content []byte

*/
func (v *pool) Dispatch(key string, content interface{}) {
	v.mMu.Lock()
	defer v.mMu.Unlock()

	_callback, _ok := v.mStore[key]
	if !_ok {
		return
	}
	_callback(content)

}

/*
Delegate
事件委派

	key string
	callback Callback

*/
func (v *pool) Delegate(key string, callback Callback) error {
	v.mMu.Lock()
	defer v.mMu.Unlock()
	if _, _ok := v.mStore[key]; _ok {
		return fmt.Errorf("duplicate key : %s", key)
	}

	v.mStore[key] = callback

	return nil
}

/*
Remove
事件移除

	key string

*/
func (v *pool) Remove(key string) {
	v.mMu.Lock()
	defer v.mMu.Unlock()

	if _, _ok := v.mStore[key]; _ok {
		v.mStore[key] = nil
		delete(v.mStore, key)
	}
}

/*
Destroy
清除變數
*/
func (v *pool) Destroy() {

	for _key := range v.mStore {
		v.Remove(_key)
	}

}

//-------------------------------------------------------------------------------------------------

type IEvent interface {
	Dispatch(key string, content interface{})
	Delegate(key string, callback Callback) error
	Remove(key string)
	Destroy()
}

//-------------------------------------------------------------------------------------------------

func New() IEvent {
	return &pool{
		mMu:    sync.Mutex{},
		mStore: map[string]Callback{},
	}
}
