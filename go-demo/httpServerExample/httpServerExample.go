package main

import (
	"fmt"
	"net/http"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платеж обработан"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}

}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello World!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно обработал HTTP запрос!")
	}

}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/pay", payHandler)

	fmt.Println("Запускаем HTTP сервер")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}
	fmt.Println("Программа закончила свое выполнение")
}
