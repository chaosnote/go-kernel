package schedule

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// FPS 120
// 1/10.9(ms)*1000(second)=92(fps)
// 1/8.3333*1000=120
const FPS = 8 * time.Millisecond

var todo = map[string]func(int64){}

// AddTask ...
func AddTask(f func(int64)) string {
	u := uuid.NewV4().String()
	todo[u] = f
	return u
}

// DelTask ...
func DelTask(key string) {
	delete(todo, key)
}

// Run ...
func Run() {

	s := time.Now().UnixNano() // start

	for range time.Tick(FPS) {
		n := time.Now().UnixNano() // now
		d := n - s                 // duration

		for _, f := range todo {
			f(d)
		}

		s = n // update start
	}

}

// After1S ...
func After1S(bt *int64, dt int64) bool {
	return AfterNS(bt, dt, 1)
}

// AfterNS ...
func AfterNS(bt *int64, dt int64, sec int64) bool {
	*bt = *bt + dt
	l := time.Duration(*bt) - time.Second*time.Duration(sec)

	if l >= 0 {
		*bt = int64(l)
		return true
	}
	return false
}

// AfterNM ...
func AfterNM(bt *int64, dt int64, min int64) bool {
	return AfterNS(bt, dt, min*60)
}

// AfterNH ...
func AfterNH(bt *int64, dt int64, hour int64) bool {
	return AfterNS(bt, dt, hour*60)
}
