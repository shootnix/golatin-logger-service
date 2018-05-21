package controllers

import (
	"encoding/json"
	"github.com/shootnix/golatin-logger-service/models"
	"log"
	"net/http"
)

type LogRequest struct {
	Module      string `json:"module"`
	Action      string `json:"action"`
	Result      string `json:"result"`
	Description string `json:"desc"`
}

type LogResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func POSTLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("Got Request!")

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	defer r.Body.Close()

	var req LogRequest
	var res LogResponse

	if err := decoder.Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res.Success = false
		res.Message = err.Error()

		encoder.Encode(res)

		return
	}

	l := models.NewLog()
	l.Module = req.Module
	l.Action = req.Action
	l.Result = req.Result
	l.Description = req.Description

	if err := l.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res.Success = false
		res.Message = err.Error()

		encoder.Encode(res)

		return
	}

	if err := l.Save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res.Success = false
		res.Message = err.Error()

		encoder.Encode(res)

		return
	}

	res.Success = true
	res.Message = "Done."

	encoder.Encode(res)
}
