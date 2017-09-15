package headAndTail

import (
	"context"
	"net/http"
	"strings"
)

func Parse(req *http.Request) (string, string, error) {
	var path string
	switch req.Context().Value("tail").(type) {
	case string:
		path = req.Context().Value("tail").(string)
	default:
		path = req.URL.String()
	}
	if path == "" {
		return "", "", nil
	}
	slash := strings.Index(path[1:], "/") + 1
	head := path[:slash]
	tail := path[slash:]
	if slash == 0 {
		head = tail
		tail = ""
	}
	return head, tail, nil
}

func With(req *http.Request) (*http.Request, error) {
	head, tail, _ := Parse(req)
	reqWith, _ := Put(req, head, tail)
	return reqWith, nil
}

func Put(req *http.Request, head, tail interface{}) (*http.Request, error) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "head", head)
	ctx = context.WithValue(ctx, "tail", tail)
	return req.WithContext(ctx), nil
}

func Get(req *http.Request) (string, string, error) {
	head := req.Context().Value("head").(string)
	tail := req.Context().Value("tail").(string)
	return head, tail, nil
}
