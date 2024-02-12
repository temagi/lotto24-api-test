package tests

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func makeAPIRequest(uri string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	url := fmt.Sprintf("%s/%s", BaseUrl, uri)
	log.Println(url)
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	err := fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
	}

	return resp
}
