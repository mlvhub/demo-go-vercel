package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlvhub/demo-go-vercel/web"
)

func FetchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user Info

	const q = `
	SELECT id, name, email, created_at, updated_at
	FROM users
	WHERE id = $1;
	`

	db, _ := web.OpenConn()
	defer db.Close()
	if err := db.GetContext(r.Context(), &user, q, id); err != nil {
		web.Response(w, nil, http.StatusInternalServerError)
		return
	}

	web.Response(w, user, http.StatusOK)
}
