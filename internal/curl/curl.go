package curl

import (
	"fmt"
	"strings"

	"github.com/rbretecher/go-postman-collection"
)

func FromRequest(request *postman.Request) string {
	return request.URL.Raw
}

func method(req *postman.Request) string {
	fmt.Println("method", req.Method)
	return "-X " + string(req.Method)
}

func headers(req *postman.Request) string {
	fmt.Printf("headers %#v\n", req.Header)
	var headers string
	const prefix = "-H "
	for _, header := range req.Header {
		headers += prefix + header.Key + ": " + header.Value + " "
	}
	return strings.TrimSuffix(headers, " ")
}

func body(req *postman.Request) string {
	fmt.Println("body", req.Body)
	return "-d " + req.Body.Raw
}