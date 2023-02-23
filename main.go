package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	v := make(chan int)
	_ = v
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := youtube.Client{}
		
		video, err := client.GetVideo("Obim8BYGnOE")


		if err != nil {
			log.Println(err)
			return
		}

		audioFormat := video.Formats.FindByItag(140)

		stream, _, err := client.GetStream(video, audioFormat)
		file, _ := os.Create("music.m4a")

		io.Copy(file, stream)
	})

	http.ListenAndServe(":8080", nil)
}
