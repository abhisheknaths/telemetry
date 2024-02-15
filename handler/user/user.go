package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhisheknaths/telemetry/internal/db"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

type user struct {
	ID       uint64 `json:"ID"`
	Username string `json:"Name"`
	Address  string `json:"Address,omitempty"`
}

type userDetail struct {
	ID      uint64 `json:"ID"`
	Address string `json:"Address"`
}

// GetUserHandler - handler for GET /users
func GetUserHandler(db db.DB, ext string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("Get user handler").Start(r.Context(), "get user task")
		defer span.End()
		rows, err := db.FetchDataRows(ctx, "SELECT * FROM telmetry_schema.telmetry_users")
		if err != nil {
			span.RecordError(err)
			return
		}
		defer rows.Close()
		u := user{}
		users := []user{}
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Username)
			if err != nil {
				span.RecordError(err)
				return
			}
			users = append(users, u)
		}
		if rows.Err() != nil {
			span.RecordError(err)
			return
		}
		span.AddEvent("Simulate delay 100ms")
		time.Sleep(time.Millisecond * 100)

		// call external service
		req, err := http.NewRequestWithContext(ctx, "GET", ext, nil)
		if err != nil {
			span.RecordError(err)
			return
		}
		// creating the client
		client := &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
		}
		res, err := client.Do(req)
		if err != nil {
			span.RecordError(err)
			return
		}
		defer res.Body.Close()
		details := &[]userDetail{}
		err = json.NewDecoder(res.Body).Decode(details)
		if err != nil {
			span.RecordError(err)
			return
		}

		detsMap := make(map[uint64]userDetail, len(users))

		for _, d := range *details {
			detsMap[d.ID] = d
		}

		for i, u := range users {
			a, ok := detsMap[u.ID]
			if ok {
				users[i].Address = a.Address
			}
		}
		bytes, _ := json.Marshal(users)
		w.Write(bytes)
	}
}

// GetUserDetail - handler for GET /users/detail
func GetUserDetail(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("Get user detail handler").Start(r.Context(), "get user detail task")
		defer span.End()
		rows, err := db.FetchDataRows(ctx, "SELECT * FROM telmetry_schema.telmetry_user_detail")
		if err != nil {
			span.RecordError(err)
			return
		}
		defer rows.Close()
		u := userDetail{}
		result := []userDetail{}
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Address)
			if err != nil {
				span.RecordError(err)
				return
			}
			result = append(result, u)
		}
		if rows.Err() != nil {
			span.RecordError(rows.Err())
			return
		}
		span.AddEvent("Simulate delay 100ms")
		time.Sleep(time.Millisecond * 100)
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}
}
