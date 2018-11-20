package userapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learn-sql/user"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserService interface {
	FindByID(id int) (*user.User, error)
}

type Handler struct {
	userService UserService
}

func writeError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, "users: "+err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}
func writeJSON(w http.ResponseWriter, value interface{}) bool {
	b, err := json.Marshal(value)
	if writeError(w, err) {
		return true
	}
	fmt.Fprintf(w, "%s", b)
	return false
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if writeError(w, err) {
		return
	}
	user, err := h.userService.FindByID(id)
	if writeError(w, err) {
		return
	}
	writeJSON(w, user)
}

func StartServer(db *sql.DB) error {
	r := mux.NewRouter()
	h := &Handler{
		userService: &user.Service{
			DB: db,
		},
	}
	r.HandleFunc("/users/{id}", h.getUser).Methods("GET")
	return http.ListenAndServe(":8000", r)
}
