testpoller - Simple polling for tests.
======================================

[![Build Status](https://travis-ci.org/LK4D4/testpoller.svg?branch=master)](https://travis-ci.org/LK4D4/testpoller)
[![GoDoc](https://godoc.org/github.com/LK4D4/testpoller?status.svg)](https://godoc.org/github.com/LK4D4/testpoller)

Sometimes you're not sure when the event will happen, using just big sleep is
one way, but it needs to be big enough to satisfy even slow machines. You can
use "polling" for speeding up your tests.

*testpoller* calls some func repeatedly until it returns true or error. It also
consumes context, which can be used to control polling time.

It's very easy to use:

```go
p := testpoller.New().WithInterval(500 * time.Millisecond)
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
err := p.Poll(ctx, func() (bool, error) {
	err := cluster.IsLeader()
	if err != nil {
		if isNotLeader(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
})
```
