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
// ndt -> now duration time
func After1S(ndt *int64, dt int64) bool {
	return AfterNS(ndt, dt, 1)
}

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

// AfterNM ...
func AfterNM(ndt *int64, dt int64, min int64) bool {
	return AfterNS(ndt, dt, min*60)
}

// AfterNH ...
func AfterNH(ndt *int64, dt int64, hour int64) bool {
	return AfterNS(ndt, dt, hour*60)
}
