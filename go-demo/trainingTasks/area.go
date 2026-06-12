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
	printRectangleArea(6, 2)
	printRectanglePerimeter(7, 3)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printCircleAreaV1(radius int) {
	circleArea, err := calculateCircleAreaV1(radius)
	check(err)

	fmt.Printf("радиус круга : %d см.\n", radius)
	fmt.Println("формула для расчета площади круга: A = π * r²")
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
	fmt.Printf("стороны треугольника : %d см, %d см, %d см \n ", a, b, c)
	fmt.Println("формула для расчета площади треугольника: S=sqrt(p*(p-a)(p-b)(p-c))")
	fmt.Printf("площадь треугольника : %d см. кв. \n", triangleArea)
}

func printTrianglePerimeter(a int, b int, c int) {
	perimeter, err := calculateTrianglePerimeter(a, b, c)
	check(err)
	fmt.Printf("стороны треугольника: %d, %d, %d\n", a, b, c)
	fmt.Println("формула периметра: P = a + b + c")
	fmt.Printf("периметр треугольника: %d см\n", perimeter)
}

func calculateTrianglePerimeter(a, b, c int) (int, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return 0, errors.New("стороны не могут быть отрицательными")
	}
	
	return a + b + c, nil
}

func calculateTriangleArea(a int, b int, c int) (int, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return int(0), errors.New("стороны не могут быть отрицательными")
	}
	p := float64(a+b+c) / 2
	triangleArea := math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))
	return int(triangleArea), nil
}

func printRectangleArea(a int, b int) {
	rectangleArea, err := calculateRectangleArea(a, b)
	check(err)

	fmt.Printf("стороны прямоугольника: %d см., %d см. \n", a, b)
	fmt.Println("формула для расчета площади прямоугольника: F=a*b")
	fmt.Printf("площадь прямоугольника : %d см. кв. \n", rectangleArea)
}
func calculateRectangleArea(a int, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return int(0), errors.New("стороны не могут быть отрицательными")
	}
	return int(float64(a * b)), nil
}

func printRectanglePerimeter(a int, b int) {
	rectanglePerimeter, err := calculateRectanglePerimeter(a, b)
	check(err)

	fmt.Printf("стороны прямоугольника: %d см., %d см. \n", a, b)
	fmt.Println("формула для расчета периметра прямоугольника: F=(a+b)*2")
	fmt.Printf("периметр прямоугольник : %d см. кв. \n", rectanglePerimeter)
}
func calculateRectanglePerimeter(a int, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return int(0), errors.New("стороны не могут быть отрицательными")
	}
	return int(float64((a + b) * 2)), nil
}
