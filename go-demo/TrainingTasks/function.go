// Задача: 1 этап) Сделать расчет площади круга в зависимости от радиуса.
/* Т.е используем 3 шага: задаем несколько значений радиуса,
считаем площадь и выводим на экран->3 функции
2 этап) Оптимизируем с помощью добавления обработки ошибок*/
package main

import (
	"errors"
	"fmt"
)

const piValue = 3.1415

func main() {
	printCircleArea(2)
}

func printCircleArea(radius int) {
	circleArea, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error()) //если ошибка не пустая, то прерываем выполнение ретерном
		return
	}
	fmt.Printf("Радиус круга : %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=pir2")
	fmt.Printf("Площадь круга : %f см. кв. \n", circleArea) //площадь круга
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Радиус круга не может быть отрицательным!")
	}
	return float32(radius) * float32(radius) * piValue, nil
}
