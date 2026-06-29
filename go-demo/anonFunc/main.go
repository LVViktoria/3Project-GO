package main

import (
	"fmt"
	"time"
)

// Обычная функция
func foo() {
	for {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

// горутина
func main() {
	// Запускаем обычную функцию в горутине
	go foo()

	// Запускаем анонимную функцию в горутине
	go func() {
		for {
			fmt.Println("Anon")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Засыпаем на 5 секунд в main горутине, чтобы дать "Foo" и "Anon" горутинам поработать
	time.Sleep(1 * time.Second)
}
