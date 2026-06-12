package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Структура пользователя
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// База данных в памяти
var database = []User{}
var currentID = 1

func main() {
	// Добавим тестового пользователя для начала
	database = append(database, User{ID: currentID, Name: "Тестовый пользователь", Email: "test@test.com"})
	currentID++

	// Маршруты
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)

	fmt.Println("\n🚀 Сервер запущен!")
	fmt.Println("📋 Откройте Postman и используйте следующие запросы:\n")
	fmt.Println("GET    - http://localhost:8080/users")
	fmt.Println("GET    - http://localhost:8080/users/1")
	fmt.Println("POST   - http://localhost:8080/users")
	fmt.Println("PUT    - http://localhost:8080/users/1")
	fmt.Println("DELETE - http://localhost:8080/users/1")
	fmt.Println("\n💡 Подсказка: Для POST и PUT не забудьте вкладку Body -> raw -> JSON\n")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Обработчик всех пользователей
func handleUsers(w http.ResponseWriter, r *http.Request) {
	// GET - получить всех пользователей
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(database)
		return
	}

	// POST - создать нового пользователя
	if r.Method == "POST" {
		var newUser User

		// Читаем JSON из тела запроса
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Ошибка: неверный JSON", http.StatusBadRequest)
			return
		}

		// Проверяем, что имя и email заполнены
		if newUser.Name == "" || newUser.Email == "" {
			http.Error(w, "Ошибка: name и email обязательны", http.StatusBadRequest)
			return
		}

		// Добавляем пользователя
		newUser.ID = currentID
		currentID++
		database = append(database, newUser)

		// Возвращаем созданного пользователя
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	// Если метод не поддерживается
	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}

// Обработчик одного пользователя по ID
func handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL (например, /users/1 -> получаем "1")
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	// Ищем пользователя в базе
	var targetUser User
	var targetIndex int
	found := false

	for i, user := range database {
		if user.ID == id {
			targetUser = user
			targetIndex = i
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}

	// GET - получить одного пользователя
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(targetUser)
		return
	}

	// PUT - обновить пользователя
	if r.Method == "PUT" {
		var updatedUser User

		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, "Ошибка: неверный JSON", http.StatusBadRequest)
			return
		}

		// Обновляем поля
		database[targetIndex].Name = updatedUser.Name
		database[targetIndex].Email = updatedUser.Email

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(database[targetIndex])
		return
	}

	// DELETE - удалить пользователя
	if r.Method == "DELETE" {
		// Удаляем из слайса
		database = append(database[:targetIndex], database[targetIndex+1:]...)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}
