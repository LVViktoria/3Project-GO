package main

import "fmt"

func main() {

	//Создаем канал
	//ch := make(chan int)

	//закрываем канал
	//close(ch)

	//1.закрываем закрытый канал
	//close(ch)

	//2.читаем значение из закрытого канал
	/*v1, ok1 := <-ch
	v2, ok2 := <-ch
	v3, ok3 := <-ch
	fmt.Println("v:", v1, v2, v3)
	fmt.Println("ok:", ok1, ok2, ok3)*/

	//3.записываем значение в закрытый канал
	//ch <- 1

	/*transferPoint := make(chan int)
	//miner
	go func() {
		iterations := 3 + rand.Intn(4)
		fmt.Println("iterations:", iterations)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}
		close(transferPoint)
	}()

	coal := 0
	/*for {
		v, ok := <-transferPoint
		if !ok {
			fmt.Println("Все обработано")
			break
		}
		coal += v
		fmt.Println("coal:", coal)
	}*/
	//короткая конструкция работы с закрытым каналом
	/*for v := range transferPoint {
		coal += v
		fmt.Println("coal:", coal)

	}
	fmt.Println("Суммарное число", coal)*/

	var ch chan string = make(chan string)

	go func() {
		ch <- "hello"
	}()

	v := <-ch

	fmt.Println("v:", v)
	close(ch)
}
