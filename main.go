package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hellow, World!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error while writing http response", err.Error())
	} else {
		fmt.Println("Success!")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "New payment processed!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error while writing http response", err.Error())
	} else {
		fmt.Println("Payment succeded. Success!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Payment cancelled!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error while writing http response", err.Error())
	} else {
		fmt.Println("Payment cancelled. Success!")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Starting listening and serving")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("An error occured", err.Error())
	}
}
