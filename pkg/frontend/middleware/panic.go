package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/sirupsen/logrus"

	"github.com/jim-minter/rp/pkg/api"
)

func Panic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				if log, ok := r.Context().Value(ContextKeyLog).(*logrus.Entry); ok {
					log.Errorf("panic: %#v\n%s\n", e, string(debug.Stack()))
				}

				api.WriteError(w, http.StatusInternalServerError, api.CloudErrorCodeInternalServerError, "", "Internal server error.")
			}
		}()

		h.ServeHTTP(w, r)
	})
}
