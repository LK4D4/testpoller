package testpoller_test

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/LK4D4/testpoller"
)

func Example() {
	p := testpoller.New()
	var counter int
	f := func() (bool, error) {
		if counter < 5 {
			fmt.Println("return false")
			counter++
			return false, nil
		}
		fmt.Println("return true")
		return true, nil
	}
	fmt.Println(p.Poll(context.Background(), f))
	// Output:
	// return false
	// return false
	// return false
	// return false
	// return false
	// return true
	// <nil>
}

func ExampleWithTimeout() {
	p := testpoller.New()
	var counter int
	f := func() (bool, error) {
		counter++
		if counter < 5 {
			return false, nil
		}
		return true, nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	fmt.Println(p.Poll(ctx, f))
	// Output:
	// context deadline exceeded
}
