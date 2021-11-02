package web

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Response(w http.ResponseWriter, data interface{}, statusCode int) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

func OpenConn() (*sqlx.DB, error) {
	q := make(url.Values)
	q.Set("sslmode", "require")
	q.Set("timezone", "utc")

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
