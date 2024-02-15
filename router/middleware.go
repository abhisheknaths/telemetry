package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/abhisheknaths/telemetry/internal/responsesnoop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const tracerPackage string = "github.com/abhisheknaths/telemetry/router"

func RequestTracer() func(next http.Handler) http.Handler {
	fn := func(next http.Handler) http.Handler {
		fn1 := func(w http.ResponseWriter, r *http.Request) {
			t := otel.Tracer(tracerPackage)
			ctx, span := t.Start(r.Context(), fmt.Sprintf(`Request start %s:%s:%s`, r.Method, r.URL, time.Now().Format(time.UnixDate)))
			defer span.End()
			r = r.WithContext(ctx)
			snooper := responsesnoop.NewSnooper(w)
			defer snooper.Release()
			next.ServeHTTP(snooper.GetWriter(), r)
			span.SetAttributes(attribute.Int("Status code", snooper.GetStatus()))
		}
		return http.HandlerFunc(fn1)
	}
	return fn
}
