package main

import (
	"fmt"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	v := make(chan int)
	_ = v
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// client := youtube.Client{}

	// video, err := client.GetVideo("Obim8BYGnOE")

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// audioFormat := video.Formats.FindByItag(140)

	// stream, _, err := client.GetStream(video, audioFormat)
	// file, _ := os.Create("music.m4a")
	ffStream := ffmpeg_go.Input("music.m4a")
	// io.Copy(file, stream)
	// fmt.Print("Here")
	err := ffStream.Output("music.mp3", ffmpeg_go.KwArgs{
		"c:a": "libmp3lame",
		"q:a": "8",
	}).OverWriteOutput().ErrorToStdOut().Run()
	fmt.Println(err)
	// })

	// http.ListenAndServe(":8080", nil)
}
