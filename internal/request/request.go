package request

import (
	"fmt"

	"github.com/rbretecher/go-postman-collection"
)

func GetRequests(collection postman.Collection) []*postman.Items {
	var requests []*postman.Items
	requests = append(requests, collection.Items...)
	return requests
}

func RequestNames(requests []*postman.Items) []string {
	var names []string
	for _, request := range requests {
		if IsDir(request) {
			names = append(names, "(d)"+request.Name)
		} else {
			var method string
			if request.Request != nil {
				method = fmt.Sprint("(", string(request.Request.Method), ") ")
			}

			names = append(names, method+request.Name)
		}
	}
	return names
}

func IsDir(request *postman.Items) bool {
	return request.Request == nil && request.Items != nil
}
