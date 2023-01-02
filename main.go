package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://www.soundhelix.com/examples/mp3/SoundHelix-Song-5.mp3")
	if err != nil {
		return
	}
	reader := bufio.NewReader(res.Body)
	var v []byte = make([]byte, 32)
	f, _ := os.Create("temp.mp3")
	buf := bufio.NewWriter(f)
	for {
		_, err = reader.Read(v)

		if err == io.EOF {
			break
		}
		buf.Write(v)
	}
}
