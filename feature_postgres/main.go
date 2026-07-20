package main

import (
	"feature_postgres/simple_connection"
	"fmt"
)

func main() {
	fmt.Println("Hello")
	simple_connection.CheckConnection()
}
