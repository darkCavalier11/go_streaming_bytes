package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://rr2---sn-cvh7knzr.googlevideo.com/videoplayback?expire=1672782830&ei=jk-0Y_n7OKGV2roP9MWfoA4&ip=202.137.218.94&id=o-AMLnrNvkLkVi_8PzQ9HycfMuikZhopnx5I2C4d9mG1wE&itag=140&source=youtube&requiressl=yes&mh=ne&mm=31%2C26&mn=sn-cvh7knzr%2Csn-h5576nsy&ms=au%2Conr&mv=m&mvi=2&pl=24&initcwndbps=1281250&spc=zIddbHhAQ_acU00k5_wF_ue0qnWV_8A&vprv=1&svpuc=1&mime=audio%2Fmp4&gir=yes&clen=5441630&dur=336.178&lmt=1655178095812376&mt=1672760941&fvip=4&keepalive=yes&fexp=24001373%2C24007246&c=ANDROID&txp=6318224&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cspc%2Cvprv%2Csvpuc%2Cmime%2Cgir%2Cclen%2Cdur%2Clmt&sig=AOq0QJ8wRAIgHQitjG66Ghef4BvmJrj56N0ovJL62VLgvTkSxYDMdUUCIA8GR8O_P8sNIUuM7oV7prFKXbJ8MltfmOL_u_KPih9t&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AG3C_xAwRgIhAKrAt60kW_Pu4ed0US4JzHneo9pJuYzA3ZqYg8j2f1f6AiEApB9JiLh5exjW4A0lGopwRWYuMDOJC2C_BQ_o7R8ZxHo%3D")
	if err != nil {
		return
	}
	reader := bufio.NewReader(res.Body)
	var v []byte = make([]byte, 32)
	f, _ := os.Create("audio.m4a")
	buf := bufio.NewWriter(f)
	for {
		_, err = reader.Read(v)

		if err == io.EOF {
			break
		}
		buf.Write(v)
	}
}

