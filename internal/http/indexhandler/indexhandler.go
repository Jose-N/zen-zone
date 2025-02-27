package indexhandler

import (
	"context"
	"net/http"

	"github.com/Jose-N/zen-zone/web/template/index"
)

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		component := index.Index()
		component.Render(context.Background(), w)
	})
}
