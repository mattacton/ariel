package curl

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rbretecher/go-postman-collection"
)

const (
	urlEncodedCT = "-H 'Content-Type: application/x-www-form-urlencoded'"
	formDataCT   = "-H 'Content-Type: multipart/form-data'"
	jsonCT       = "-H 'Content-Type: application/json'"
)

func FromRequest(request *postman.Request) string {
	var s strings.Builder
	s.WriteString("curl ")
	s.WriteString(method(request))
	s.WriteString(" ")
	s.WriteString(headers(request))
	s.WriteString(" ")
	s.WriteString(body(request))
	s.WriteString(" ")
	s.WriteString(request.URL.Raw)
	return s.String()
}

func method(req *postman.Request) string {
	return "-X " + string(req.Method)
}

func headers(req *postman.Request) string {
	var headers string
	const prefix = "-H "
	for _, header := range req.Header {
		headers += prefix + header.Key + ": " + header.Value + " "
	}
	return strings.TrimSuffix(headers, " ")
}

func body(req *postman.Request) string {
	if req.Body != nil {
		switch req.Body.Mode {
		case "raw":
			var ct string
			if isJSON(req.Body) {
				ct = jsonCT + " "
			}
			return fmt.Sprint(ct, rawBody(req.Body))
		case "urlencoded":
			return fmt.Sprint(urlEncodedCT, " ", urlEncodedBody(req.Body))
		case "formdata":
			return fmt.Sprint(formDataCT, " ", formDataBody(req.Body))
		// case "file":
		// case "graphql":
		default:
			return "--data '" + strings.ReplaceAll(req.Body.Raw, "\n", "") + "'"
		}
	}
	return ""
}

func formDataBody(body *postman.Body) string {
	if body.FormData != nil {
		var bodyStrs []string
		if maps, ok := body.FormData.([]interface{}); ok {
			for _, mapz := range maps {
				if mapt, ok := mapz.(map[string]interface{}); ok {
					fmt.Println("keyval", mapt["key"], mapt["value"])
					key := url.QueryEscape(mapt["key"].(string))
					val := url.QueryEscape(mapt["value"].(string))
					bodyStrs = append(bodyStrs, key+"="+val)
				}
			}
		}
		return "--data '" + strings.Join(bodyStrs, "&") + "'"
	}
	return ""
}

func urlEncodedBody(body *postman.Body) string {
	if body.URLEncoded != nil {
		var bodyStrs []string
		if maps, ok := body.URLEncoded.([]interface{}); ok {
			for _, mapz := range maps {
				if mapt, ok := mapz.(map[string]interface{}); ok {
					fmt.Println("keyval", mapt["key"], mapt["value"])
					key := url.QueryEscape(mapt["key"].(string))
					val := url.QueryEscape(mapt["value"].(string))
					bodyStrs = append(bodyStrs, key+"="+val)
				}
			}
		}
		return "--data '" + strings.Join(bodyStrs, "&") + "'"
	}
	return ""
}

func rawBody(body *postman.Body) string {
	return "--data '" + strings.ReplaceAll(body.Raw, "\n", "") + "'"
}

// yes this is simplistic, but it's a start
func isJSON(body *postman.Body) bool {
	if body == nil {
		return false
	}
	return strings.Contains(body.Raw, "{") && strings.Contains(body.Raw, "}")
}
