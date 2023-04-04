package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"net/http"
	"strconv"
)

func (h *Manager) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	var user model.UserCreateReq
	json.NewDecoder(r.Body).Decode(&user)
	res := h.srv.User.Update(id, user)
	if res != nil {
		http.Error(w, fmt.Sprintf("%s", res), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{Message: "User updated successfully"})
}

func (h *Manager) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	err = h.srv.User.Delete(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{Message: "User deleted successfully"})
}

func (h *Manager) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	user, err := h.srv.User.GetById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
