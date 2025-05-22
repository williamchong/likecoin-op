package middleware

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
)

func MakeSentryMiddleware(sentryHub *sentry.Hub) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hub := sentryHub.Clone()
			hub.Scope().SetRequest(r)
			r = r.WithContext(sentry.SetHubOnContext(r.Context(), hub))

			defer func() {
				v := recover()
				if v != nil {
					hub.RecoverWithContext(r.Context(), v)
					sentry.Flush(time.Second * 5)
					panic(v)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
