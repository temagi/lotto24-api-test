package tests

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func makeAPIRequest(baseURL string, uri string) *fasthttp.Response {
	req := fasthttp.AcquireRequest()
	url := fmt.Sprintf("%s/%s", baseURL, uri)
	log.Println(url)
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	fasthttp.Do(req, resp)

	return resp
}