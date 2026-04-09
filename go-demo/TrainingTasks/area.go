/*Задача: перепишите программу для подсчета площади круга, используя число Пи из
библиотеки math. Напишите функции для подсчета площади и периметра
прямоугольника и треугольника. + добавить описание ошибки*/

package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	printCircleAreaV1(5)
	printTriangleArea(3, 7, 1)
	printTrianglePerimeter(4, 1, 5)
	//printRectangleArea
	//printRectanglePerimeter
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
} //доработать обработчик??

func printCircleAreaV1(radius int) {
	circleArea, err := calculateCircleAreaV1(radius)
	check(err)

	fmt.Printf("радиус круга : %d см.\n", radius)
	fmt.Println("формула для расчета площади круга: A=pir2")
	fmt.Printf("площадь круга : %f см. кв. \n", circleArea) //площадь круга
}

func calculateCircleAreaV1(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("радиус круга не может быть отрицательным")
	}
	return float32(math.Pi * math.Pow(float64(radius), 2)), nil
}
func printTriangleArea(a int, b int, c int) {
	triangleArea, err := calculateTriangleArea(a, b, c)
	check(err)
	fmt.Printf("стороны треугольника : %d см, %d, см, %d см \n ", a, b, c)
	fmt.Println("формула для расчета площади треугольника: S=sqrt(p*(p-a)(p-b)(p-c))")
	fmt.Printf("площадь треугольника : %d см. кв. \n", triangleArea)
}

func printTrianglePerimeter(a int, b int, c int) (int, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return int(0), errors.New("стороны не могут быть отрицательными")
	}
	p := a + b + c
	return p, nil
}

func calculateTriangleArea(a int, b int, c int) (int, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return int(0), errors.New("стороны не могут быть отрицательными")
	}
	p := a + b + c
	return int(math.Sqrt(float64(p / 2 * (p/2 - a) * (p/2 - b) * (p/2 - c)))), nil
}
