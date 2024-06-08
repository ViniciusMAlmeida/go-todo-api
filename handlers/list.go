package handlers

import (
	"encoding/json"
	"go-todo-api/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
