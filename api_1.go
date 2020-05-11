package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request recieved!")
		fmt.Println("Method: ", r.Method)
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request recieved!")
		fmt.Println("Method: ", r.Method)
		w.Write([]byte("Anything after get going!"))
	})
	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request recieved!")
		fmt.Println("Method: ", r.Method)
		w.Write([]byte("go only!"))
	})
	http.ListenAndServe("localhost:3001", mux)
	fmt.Println("Listening on: localhost:3001/getgoing") //curl -i -X POST localhost:3001/getgoing/

}
