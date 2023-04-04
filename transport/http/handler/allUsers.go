package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/model"
	"log"
	"net/http"
)

func (h *Manager) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.UserCreateReq
	var resp model.UserCreateResp
	json.NewDecoder(r.Body).Decode(&user)
	resp, err := h.srv.AllUsers.Create(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resp)

}

func (h *Manager) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.srv.AllUsers.Get()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
