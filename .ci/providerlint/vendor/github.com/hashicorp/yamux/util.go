package yamuximport (
	"sync"
	"time"
)var (
	timerPool = &sync.Pool{
		New: 
() interface{} {
			timer := time.NewTimer(time.Hour * 1e6)
			timer.Stop()
			return timer
		},
	}
)syncSendErr is used to try an async send of an error asyncSendErr(ch chan error, err error) {
	if ch == nil {
		return
	}
	select {
	case ch <- err:
	default:
	}
}// asyncNotify is used to signal a waiting goroutine asyncNotify(ch chan struct{}) {
	select {
	case ch <- struct{}{}:
	default:
	}
// min computes the minimum of two values min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
