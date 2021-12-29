package schedule

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

// FPS 120
// 1/10.9(ms)*1000(second)=92(fps)
// 1/8.3333*1000=120
const FPS = 8 * time.Millisecond

//-------------------------------------------------------------------------------------------------

type task struct {
	mu    sync.Mutex
	store map[string]func(int64)
}

/*
AddTask ...
*/
func (v *task) AddTask(f func(int64)) string {
	v.mu.Lock()
	defer v.mu.Unlock()

	u := uuid.NewV4().String()
	v.store[u] = f
	return u
}

/*
DelTask ...
*/
func (v *task) DelTask(key string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	delete(v.store, key)
}

/*
Run ...
*/
func (v *task) Run(d int64) {
	v.mu.Lock()
	defer v.mu.Unlock()

	for _, f := range v.store {
		f(d)
	}
}

//-------------------------------------------------------------------------------------------------

var _task = &task{
	store: map[string]func(int64){},
}

// AddTask ...
func AddTask(f func(int64)) string {
	return _task.AddTask(f)
}

// DelTask ...
func DelTask(key string) {
	_task.DelTask(key)
}

//-------------------------------------------------------------------------------------------------

func init() {

	go func() {

		s := time.Now().UnixNano() // start

		for range time.Tick(FPS) {
			n := time.Now().UnixNano() // now
			d := n - s                 // duration

			_task.Run(d)

			s = n // update start
		}

	}()

}

//-------------------------------------------------------------------------------------------------

// AfterNS ...
func AfterNS(ndt *int64, dt int64, sec int64) bool {
	*ndt = *ndt + dt
	l := time.Duration(*ndt) - time.Second*time.Duration(sec)

	if l >= 0 {
		*ndt = int64(l)
		return true
	}
	return false
}

// After1S ...
// ndt -> now duration time
func After1S(ndt *int64, dt int64) bool {
	return AfterNS(ndt, dt, 1)
}

// AfterNM ...
func AfterNM(ndt *int64, dt int64, min int64) bool {
	return AfterNS(ndt, dt, min*60)
}

// AfterNH ...
func AfterNH(ndt *int64, dt int64, hour int64) bool {
	return AfterNS(ndt, dt, hour*60)
}
