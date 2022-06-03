package main

import (
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://elencos.com.br/wp-content/uploads/2021/11/inter-2010-1-1.jpeg"

func main() {
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}
	res.Body.Close()

	enc := b64.StdEncoding.EncodeToString(data)
	log.Println(enc)

	//ioutil.WriteFile("google_logo.png", data, 0666)

	log.Println("Sucesso ao converter a imagem!")
}
