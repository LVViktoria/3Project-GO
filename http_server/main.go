package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	str := "message:pong"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Обработка на паттерне /ping")
	}
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	str := "status:ok"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Обработка на паттерне /health")
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	str := "version:1.0.0"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Обработка на паттерне /version")
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	fmt.Println("Запускаем HTTP сервер")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}
	fmt.Println("Программа закончила свое выполнение")
}
