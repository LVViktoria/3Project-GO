package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo завершилась", n)
			return
		default:
			fmt.Println("Foo продолжает свое выполнение", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
func boo(ctx context.Context, n int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo завершилась", n)
			return
		default:
			fmt.Println("Boo продолжает свое выполнение", n)
		}

		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)
	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)

	go boo(childContext, 1)
	go boo(childContext, 2)
	go boo(childContext, 3)

	time.Sleep(1 * time.Second)
	childCancel()

	time.Sleep(1 * time.Second)
	parentCancel()

	time.Sleep(3 * time.Second)
	fmt.Println("main завершился")
}
