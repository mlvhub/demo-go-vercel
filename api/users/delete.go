package users

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlvhub/demo-go-vercel/web"
)

func Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	const q = `DELETE FROM users WHERE id = $1;`

	db, _ := web.OpenConn()
	defer db.Close()
	if _, err := db.ExecContext(r.Context(), q, id); err != nil {
		web.Response(w, nil, http.StatusInternalServerError)
		return
	}

	web.Response(w, nil, http.StatusOK)
}
