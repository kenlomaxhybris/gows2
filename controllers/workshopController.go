package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kenlomaxhybris/goworkshopII/models"
)

var repo models.WorkshopRepo

const MAX = 100

func extractWorkshopFromPayload(r *http.Request) models.Workshop {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ws models.Workshop
	json.Unmarshal(reqBody, &ws)
	return ws
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CreateWorkshop(w http.ResponseWriter, r *http.Request) {
	ws := extractWorkshopFromPayload(r)
	if len(ws.Presenter) == 0 || len(ws.Presenter) > MAX || len(ws.Title) == 0 || len(ws.Title) > MAX {
		respondWithError(w, http.StatusUnprocessableEntity, "Missing/too much Data")
		return
	}
	ws = repo.Create(ws)
	respondWithJSON(w, http.StatusOK, ws)
}

func ReadWorkshop(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ws, e := repo.Read(id)
	if e != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("ID %d not found", id))
		return
	}
	respondWithJSON(w, http.StatusOK, ws)
}

func ReadAllWorkshops(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, repo.ReadAll())
}

func UpdateWorkshop(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	ws := extractWorkshopFromPayload(r)
	if len(ws.Presenter) == 0 || len(ws.Presenter) > MAX || len(ws.Title) == 0 || len(ws.Title) > MAX {
		respondWithError(w, http.StatusUnprocessableEntity, "Missing/too much Data")
		return
	}

	ws, e := repo.Update(ws, id)
	if e != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("ID %d not found", id))
		return
	}
	respondWithJSON(w, http.StatusOK, ws)
}

func DeleteWorkshop(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	e := repo.Delete(id)
	if e != nil {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("ID %d not found", id))
		return
	}
	respondWithJSON(w, http.StatusOK, "")
}
