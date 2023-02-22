package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		fmt.Println(values.Get("videoId"))
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)
}
