package db

import (
	"context"
	"fmt"

	"github.com/abhisheknaths/telemetry/internal/instrumentation"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

type rows struct {
	r pgx.Rows
}

const tracerPackageName string = "github.com/abhisheknaths/telemetry/internal/db"

func (db *db) FetchDataRows(ctx context.Context, query string, args ...any) (DataRows, error) {
	r, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	rows := new(rows)
	rows.r = r
	return r, nil
}

func (r *rows) Scan(dest ...any) error {
	return r.Scan(dest...)
}

func (r *rows) Next() bool {
	return r.r.Next()
}

func (r *rows) Err() error {
	return r.Err()
}

func (r *rows) Close() {
	r.Close()
}

func NewDB(connString string, tp instrumentation.TracerProvider) (DB, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	withTraceProvider := otelpgx.WithTracerProvider(tp)
	config.ConnConfig.Tracer = otelpgx.NewTracer(withTraceProvider)

	pool, err := pgxpool.NewWithConfig(context.TODO(), config)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.TODO())
	if err != nil {
		return nil, fmt.Errorf(`ping test failed with error %w`, err)
	}
	return &db{
		pool: pool,
	}, nil
}
