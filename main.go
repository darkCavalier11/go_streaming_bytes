package main

import (
	"log"
	"net/http"

	"github.com/kkdai/youtube/v2"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	v := make(chan int)
	_ = v
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client := youtube.Client{}
		video, err := client.GetVideo("Obim8BYGnOE")
		audioFormat := video.Formats.FindByItag(140)
		url, err := client.GetStreamURL(video, audioFormat)

		if err != nil {
			log.Println(err)
			return
		}
		ffStream := ffmpeg_go.Input(url)
		err = ffStream.Output("music.mp3", ffmpeg_go.KwArgs{
			"acodec": "libmp3lame",
		}).OverWriteOutput().Run()
	})

	http.ListenAndServe(":8080", nil)
}
