package testpoller

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestPollerOK(t *testing.T) {
	p := New()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := p.Poll(ctx, func() (bool, error) { return true, nil })
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPollerError(t *testing.T) {
	p := New()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	testErr := errors.New("test error")
	err := p.Poll(ctx, func() (bool, error) { return false, testErr })
	if err == nil {
		t.Fatal("error must be spawned")
	}
	if err != testErr {
		t.Fatalf("got error: %v, expected %v", err, testErr)
	}
}

func TestPollerCancelled(t *testing.T) {
	p := New()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := p.Poll(ctx, func() (bool, error) { return false, nil })
	if err == nil {
		t.Fatal("error must be spawned")
	}
	if err != context.Canceled {
		t.Fatalf("got error: %v, expected %v", err, context.Canceled)
	}
}

func TestPollerTimeout(t *testing.T) {
	p := New()
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := p.Poll(ctx, func() (bool, error) { return false, nil })
	if err == nil {
		t.Fatal("error must be spawned")
	}
	if err != context.DeadlineExceeded {
		t.Fatalf("got error: %v, expected %v", err, context.DeadlineExceeded)
	}
}

func TestPollerCancelLater(t *testing.T) {
	p := New()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()
	err := p.Poll(ctx, func() (bool, error) { return false, nil })
	if err == nil {
		t.Fatal("error must be spawned")
	}
	if err != context.Canceled {
		t.Fatalf("got error: %v, expected %v", err, context.Canceled)
	}
}

func TestPollerGotTrue(t *testing.T) {
	p := New()
	var cnt int
	err := p.Poll(context.Background(), func() (bool, error) {
		cnt++
		if cnt == 5 {
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
