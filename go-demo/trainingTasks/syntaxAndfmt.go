// Задача: изучить синтаксис и возможности пакета fmt
package main

import "fmt"

const pi = 3.1415

func main() {

	circleRadius := 2
	circleArea := float32(circleRadius) * float32(circleRadius) * pi

	fmt.Printf("Радиус круга : %d см.\n", circleRadius)
	fmt.Println("Формула для расчета площади круга: A=pir2")

	fmt.Printf("Площадь круга : %f см. кв. \n", circleArea)

	/*userHeight := 1.8
	var userKg
	userKg = 100
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)*/
}
