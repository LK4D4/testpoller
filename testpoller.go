package testpoller

import (
	"sync"
	"time"

	"golang.org/x/net/context"
)

const defaultInterval = 100 * time.Millisecond

// Poller is configuration struct for Poll. It stores polling interval, which by
// default is 100 milliseconds.
type Poller struct {
	interval time.Duration
	once     sync.Once
}

// New returns new Poller with interval set to 100 milliseconds.
func New() Poller {
	return Poller{
		interval: 100 * time.Millisecond,
	}
}

// WithInterval returns new Poller with interval set to d.
func (p Poller) WithInterval(d time.Duration) Poller {
	return Poller{
		interval: d,
	}
}

// Poll executes f with interval until it returns true or error. It returns
// error if f returns error or ctx is cancelled.
func (p Poller) Poll(ctx context.Context, f func() (bool, error)) error {
	p.once.Do(func() {
		if p.interval == 0 {
			p.interval = defaultInterval
		}
	})
	res, err := f()
	if err != nil {
		return err
	}
	if res {
		return nil
	}
	timer := time.NewTicker(p.interval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			res, err := f()
			if err != nil {
				return err
			}
			if res {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}

	}
}
