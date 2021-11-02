package users

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mlvhub/demo-go-vercel/web"
)

type NewUser struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user NewUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		web.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = uuid.New().String()

	const q = `
		INSERT INTO users (id, name, email)
		VALUES ($1, $2, $3);
	`

	db, _ := web.OpenConn()
	defer db.Close()
	if _, err := db.ExecContext(r.Context(), q, user.ID, user.Name, user.Email); err != nil {
		web.Response(w, err.Error(), http.StatusInternalServerError)
		return
	}

	web.Response(w, user, http.StatusCreated)
}
