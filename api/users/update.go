package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlvhub/demo-go-vercel/web"
)

type UpdateUser struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user UpdateUser

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		web.Response(w, nil, http.StatusInternalServerError)
		return
	}

	const q = `
		UPDATE users
		SET name = $1, email = $2, updated_at = NOW()
		WHERE id = $3;
	`

	db, _ := web.OpenConn()
	defer db.Close()
	if _, err := db.ExecContext(r.Context(), q, user.Name, user.Email, id); err != nil {
		web.Response(w, nil, http.StatusInternalServerError)
		return
	}

	web.Response(w, user, http.StatusOK)
}
