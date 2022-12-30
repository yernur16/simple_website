package router

import (
	"fmt"
	"net/http"
	"strings"
)

func extractTaskId(r *http.Request) (string, error) {
	if len(r.URL.Path) < 2 {
		return "", fmt.Errorf("no task id provided")
	}
	// We know our subrouter will only provide the id as the root path item.
	return strings.Split(r.URL.Path[1:], "/")[0], nil
}
