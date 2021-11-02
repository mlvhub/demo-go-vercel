package users

import (
	"net/http"
	"time"

	"github.com/mlvhub/demo-go-vercel/web"
)

type Info struct {
	ID        string    `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name,omitempty"`
	Email     string    `db:"email" json:"email,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

func FetchAll(w http.ResponseWriter, r *http.Request) {
	const q = `
		SELECT id, name, email, created_at, updated_at
		FROM users
		ORDER BY created_at DESC;
	`
	var users []Info

	db, err := web.OpenConn()
	if err != nil {
		web.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	if err := db.SelectContext(r.Context(), &users, q); err != nil {
		web.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		users = make([]Info, 0)
	}

	web.Response(w, users, http.StatusOK)
}
