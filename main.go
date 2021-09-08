package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("hi")
		resp.Write([]byte("hello"))
	})
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
