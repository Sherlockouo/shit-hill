package util

import (
	"gin-toy/constant"
	"log"
	"net/http"
)

var (
	cli *http.Client
)

func Do(req *http.Request) (*http.Response, error) {
	log.Default().Println(constant.BaseURL)
	defer func() {
		if r := recover(); r != nil {
			log.Panicln(r)
		}
	}()
	cli = &http.Client{}
	return cli.Do(req)
}

func GET(url string) (*http.Response, error) {
	cli = &http.Client{}
	return cli.Get(url)
}
