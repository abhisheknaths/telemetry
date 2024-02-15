package router

import (
	"net/http"

	handler "github.com/abhisheknaths/telemetry/handler/user"
	"github.com/abhisheknaths/telemetry/internal/db"
	"github.com/abhisheknaths/telemetry/internal/instrumentation"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type RouterDeps struct {
	Database      db.DB
	Version       string
	TraceProvider instrumentation.TracerProvider
}

type RouterDepsWithExternalDeps struct {
	RouterDeps
	ExternalURL string
}

func InitRouter(ed RouterDepsWithExternalDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(RequestTracer())
	r.Get("/users", handler.GetUserHandler(ed.Database, ed.ExternalURL))
	return r
}

func InitDetailRouter(rd RouterDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(RequestTracer())
	r.Get("/user/detail", handler.GetUserDetail(rd.Database))
	// wrapping the chi mux handler with instrumentation from the otelhttp library
	ih := otelhttp.NewHandler(r, "app2-handler-wrap")
	return ih
}
