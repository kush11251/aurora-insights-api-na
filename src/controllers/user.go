package controllers

import (
	"aurora-insights-api/src/models"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := models.GetUser(id, db)
	json.NewEncoder(w).Encode(user)
}

func RegisterControllers(r *mux.Router) {
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
}
