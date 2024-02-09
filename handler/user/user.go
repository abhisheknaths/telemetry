package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abhisheknaths/telemetry/internal/db"
	"go.opentelemetry.io/otel"
)

type user struct {
	ID       uint64 `json:"ID"`
	Username string `json:"Name"`
}

// GetUserHandler - handler for GET /users
func GetUserHandler(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("Get user handler").Start(r.Context(), "get user task")
		defer span.End()
		rows, err := db.FetchDataRows(ctx, "SELECT * FROM telmetry_schema.telmetry_users")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		u := user{}
		result := []user{}
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Username)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, u)
		}
		span.AddEvent("Simulate delay 100ms")
		time.Sleep(time.Millisecond * 100)
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}
}

// GetUserDetail - handler for GET /users/1/detail
func GetUserDetail(db db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := otel.Tracer("Get user handler").Start(r.Context(), "get user task")
		defer span.End()
		rows, err := db.FetchDataRows(ctx, "SELECT * FROM telmetry_schema.telmetry_users")
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		u := user{}
		result := []user{}
		for rows.Next() {
			err = rows.Scan(&u.ID, &u.Username)
			if err != nil {
				fmt.Println(err)
			}
			result = append(result, u)
		}
		span.AddEvent("Simulate delay 100ms")
		time.Sleep(time.Millisecond * 100)
		bytes, _ := json.Marshal(result)
		w.Write(bytes)
	}
}
