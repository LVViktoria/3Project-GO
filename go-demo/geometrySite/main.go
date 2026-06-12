package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Структуры для формул
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

type Rectangle struct {
	Length, Width float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Triangle struct {
	SideA, SideB, SideC float64
	Base, Height        float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// Вспомогательная функция для парсинга float
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// Главный обработчик
func indexHandler(w http.ResponseWriter, r *http.Request) {
	var result string
	var calcType string
	var calcData map[string]float64

	if r.Method == "POST" {
		figure := r.FormValue("figure")
		calcData = make(map[string]float64)

		switch figure {
		case "circle":
			radius := parseFloat(r.FormValue("radius"))
			circle := Circle{Radius: radius}
			area := circle.Area()
			result = fmt.Sprintf("Площадь круга: %.2f кв. ед.", area)
			calcType = "circle"
			calcData["radius"] = radius
			calcData["result"] = area

		case "rectangle_area":
			length := parseFloat(r.FormValue("length"))
			width := parseFloat(r.FormValue("width"))
			rect := Rectangle{Length: length, Width: width}
			area := rect.Area()
			result = fmt.Sprintf("Площадь прямоугольника: %.2f кв. ед.", area)
			calcType = "rectangle_area"
			calcData["length"] = length
			calcData["width"] = width
			calcData["result"] = area

		case "rectangle_perimeter":
			length := parseFloat(r.FormValue("length"))
			width := parseFloat(r.FormValue("width"))
			rect := Rectangle{Length: length, Width: width}
			perimeter := rect.Perimeter()
			result = fmt.Sprintf("Периметр прямоугольника: %.2f ед.", perimeter)
			calcType = "rectangle_perimeter"
			calcData["length"] = length
			calcData["width"] = width
			calcData["result"] = perimeter

		case "triangle_area":
			base := parseFloat(r.FormValue("base"))
			height := parseFloat(r.FormValue("height"))
			tri := Triangle{Base: base, Height: height}
			area := tri.Area()
			result = fmt.Sprintf("Площадь треугольника: %.2f кв. ед.", area)
			calcType = "triangle_area"
			calcData["base"] = base
			calcData["height"] = height
			calcData["result"] = area

		case "triangle_perimeter":
			sideA := parseFloat(r.FormValue("sideA"))
			sideB := parseFloat(r.FormValue("sideB"))
			sideC := parseFloat(r.FormValue("sideC"))
			tri := Triangle{SideA: sideA, SideB: sideB, SideC: sideC}
			perimeter := tri.Perimeter()
			result = fmt.Sprintf("Периметр треугольника: %.2f ед.", perimeter)
			calcType = "triangle_perimeter"
			calcData["sideA"] = sideA
			calcData["sideB"] = sideB
			calcData["sideC"] = sideC
			calcData["result"] = perimeter
		}

		// Отдаём результат вместе с данными для сохранения в localStorage
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		data := struct {
			Result     string
			CalcType   string
			CalcData   map[string]float64
			ShowResult bool
		}{
			Result:     result,
			CalcType:   calcType,
			CalcData:   calcData,
			ShowResult: true,
		}
		tmpl.Execute(w, data)
		return
	}

	// GET запрос — просто показываем форму
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
