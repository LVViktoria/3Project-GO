package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Payment struct {
	Description string `json:"description"` //Описание покупки
	USD         int    `json:"usd"`         //Сумма покупки
	FullName    string `json:"fullName"`    //Фио человека совершающего покупки
	Address     string `json:"address"`     //Место прописки человека совершающего покупки
}

func (p Payment) Println() {
	fmt.Println("Description: ", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
}

var mtx = sync.Mutex{}
var money = 1000
var paymentHistory = make([]Payment, 0)

func payHandler(w http.ResponseWriter, r *http.Request) {

	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// или как ниже
	/*httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var payment Payment
	if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	payment.Println()

	/*httpRequestBodyString := string(httpRequestBody)
	parts := strings.SplitN(httpRequestBodyString, ",", 2)
	if len(parts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
	}

	usd, err := strconv.Atoi(parts[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payment := Payment{
		Description: parts[1],
		USD:         usd,
	}*/
	mtx.Lock()
	if money-payment.USD >= 0 {
		money -= payment.USD
	}
	paymentHistory = append(paymentHistory, payment)
	fmt.Println("money:", money)
	fmt.Println("payment history:", paymentHistory)
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)
	/*err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка во время работы http сервера", err)
	}*/

	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы http сервера", err) //почему здесь переменная err исчезает после условного ветвления?
	}
}
