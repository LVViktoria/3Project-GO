package main

import "fmt"

func main() {
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			fmt.Printf("i = %d, j = %d\n", i, j)

			if i == 2 && j == 2 {

				fmt.Println("Выход из внешнего цикла...")

				break OuterLoop
			}
		}
	}

	fmt.Println("Цикл завершен...")
}
