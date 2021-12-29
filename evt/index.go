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
	mQueue bool
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

	if v.mQueue {
		_callback(content)
	} else {
		go _callback(content)
	}

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
	v.mMu.Lock()
	defer v.mMu.Unlock()

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

var mEvent = New(true)

func Dispatch(key string, content interface{}) {
	mEvent.Dispatch(key, content)
}

func Delegate(key string, callback Callback) error {
	return mEvent.Delegate(key, callback)
}

func Remove(key string) {
	mEvent.Remove(key)
}

func Destroy() {
	mEvent.Destroy()
}

//-------------------------------------------------------------------------------------------------

/*
New

	queue 是否使用隊列
		false 則使用 go func()

*/
func New(queue bool) IEvent {
	return &pool{
		mMu:    sync.Mutex{},
		mQueue: queue,
		mStore: map[string]Callback{},
	}
}
