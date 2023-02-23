package main

import (
	"net/http"
)

func main() {
	v := make(chan int)
	_ = v
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		videoId := values.Get("videoId")
		_ = videoId
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", nil)
}
