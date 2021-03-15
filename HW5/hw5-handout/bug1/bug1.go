package bug1

import "sync"

// Counter stores a count.
type Counter struct {
	n int64
}


var mu sync.Mutex

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	mu.Lock()
	c.n++
	mu.Unlock()
}

