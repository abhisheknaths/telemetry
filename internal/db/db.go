package db

import "context"

type DB interface {
	FetchDataRows(ctx context.Context, query string, args ...any) (DataRows, error)
}

type DataRows interface {
	Scan(dest ...any) error
	Next() bool
	Err() error
	Close()
}
