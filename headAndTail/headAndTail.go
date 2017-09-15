package headAndTail

import (
	"context"
	"net/http"
	"strings"
)

func WithHeadAndTail(req *http.Request, head, tail interface{}) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "head", head)
	ctx = context.WithValue(ctx, "tail", tail)
	return req.WithContext(ctx)
}

func HeadAndTail(req *http.Request) (string, string) {
	var path string
	switch req.Context().Value("tail").(type) {
	case string:
		path = req.Context().Value("tail").(string)
	default:
		path = req.URL.String()
	}
	if path == "" {
		return "", ""
	}
	slash := strings.Index(path[1:], "/") + 1
	head := path[:slash]
	tail := path[slash:]
	if slash == 0 {
		head = tail
		tail = ""
	}
	return head, tail
}
