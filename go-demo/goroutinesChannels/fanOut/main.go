package main

import (
	"fmt"
	"time"
)

// Функция, описывающая поход в шахту
// n -> номер похода в шахту
func mine(transferPoint chan int, n int) int {
	// Имитируем поход в шахту на одну секунду
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	// Возвращаем добытый уголь
	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передал")
	return n
}

// горутина
func main() {
	coal := 0

	// Засекаем время
	initTime := time.Now()

	transferPoint := make(chan int)
	// Последовательно, раз за разом, ходим 3 раза в шахту
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)
	go mine(transferPoint, 4)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	// Итоговые значения угля и времени выполнения
	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))
}
