package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abhisheknaths/telemetry/internal/config"
	"github.com/abhisheknaths/telemetry/internal/db"
	"github.com/abhisheknaths/telemetry/internal/instrumentation"
	"github.com/abhisheknaths/telemetry/router"
)

var version string

const serviceName string = "fetch-user-service"

func main() {
	ctx := context.Background()
	dbConfig, err := config.LoadDBConfig()
	if err != nil {
		fmt.Println("failed to read db config")
		panic(err)
	}

	instrumentationConfig, err := config.LoadInstrumentationConfig()
	if err != nil {
		fmt.Println("failed to read instrumentation config")
		panic(err)
	}

	traceExporter, err := instrumentation.NewHTTPTraceExporter(ctx, instrumentationConfig.ExporterEndpoint, instrumentationConfig.ExporterPath)
	if err != nil {
		fmt.Println("failed to establish trace exporter connection")
		panic(err)
	}

	tp := instrumentation.NewTracerProvider(traceExporter, serviceName)
	instrumentation.SetTraceProviderGlobally(tp)

	db, err := db.NewDB(dbConfig.ConnString, tp)
	if err != nil {
		fmt.Println("failed to establish db connection")
		panic(err)
	}

	rd := router.RouterDeps{
		Database: db,
		Version:  version,
	}
	r := router.InitRouter(rd)
	fmt.Println("listening on 3000")
	http.ListenAndServe(":3000", r)
}
