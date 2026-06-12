package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var mtx = sync.Mutex{}
var money = 1000 //usd
var bank = 0     //usd

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	//fmt.Println("HTTP method", r.Method)

	/*for k, v := range r.Header {
		fmt.Println("k:", k, "--v:", v)
	}*/
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) //не смогли прочитать

		msg := "fail to read HTTP body:" + err.Error() //msg-message
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //не корректный запрос со стороны клиента
		msg := "fail to convert HTTP body to integer:" + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}

	mtx.Lock()
	if money-paymentAmount >= 0 {

		money -= paymentAmount
		msg := "Оплата прошла успешно: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP body: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "fail to convert HTTP body to integer: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
			return
		}
	}

	mtx.Lock()
	if money >= saveAmount {

		money -= saveAmount //взяли деньги из кошелька
		msg := "Денег в копилке: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}

		bank += saveAmount //положили в копилку
		msg = "Баланс кошелька: " + strconv.Itoa(bank)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to write HTTP response", err)
		}
	}
	mtx.Unlock()
}

func main() {

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP SERVER ERROR:", err)
	}

}
