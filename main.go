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
	str := "Hellow, World!"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Error while writing http response", err.Error())
	} else {
		fmt.Println("Payment succeded. Success!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Hellow, World!"
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
}
