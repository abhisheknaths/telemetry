package router

import (
	"net/http"

	handler "github.com/abhisheknaths/telemetry/handler/user"
	"github.com/abhisheknaths/telemetry/internal/db"
	"github.com/abhisheknaths/telemetry/internal/instrumentation"
	"github.com/go-chi/chi/v5"
)

type RouterDeps struct {
	Database      db.DB
	Version       string
	TraceProvider instrumentation.TracerProvider
}

func InitRouter(rd RouterDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(RequestTracer())
	r.Get("/users", handler.GetUserHandler(rd.Database))
	return r
}
