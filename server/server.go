package server

import (
	"drone/drone"
	"drone/medication"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func (h handler) GetDrones(w http.ResponseWriter, r *http.Request) {
	drones, err := drone.GetDrones(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(drones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (h handler) RegisterDrone(w http.ResponseWriter, r *http.Request) {
	var newDrone drone.Drone

	err := json.NewDecoder(r.Body).Decode(&newDrone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	drones, err := drone.RegisterDrone(h.DB, newDrone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(drones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (h handler) CheckingLoadedMedication(w http.ResponseWriter, r *http.Request) {
	droneID, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, "Couldn't find id in request URL")))
		return
	}
	id, err := strconv.Atoi(droneID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, "Invaild drone id")))
		return
	}
	droneMedications, err := drone.CheckingLoadedMedication(h.DB, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(droneMedications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (h handler) LoadMedication(w http.ResponseWriter, r *http.Request) {
	droneID, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, "Couldn't find id in request URL")))
		return
	}
	id, err := strconv.Atoi(droneID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, "Invaild drone id")))
		return
	}
	var medications []medication.Medication
	err = json.NewDecoder(r.Body).Decode(&medications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = drone.LoadMedication(h.DB, medications, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

func (h handler) AvailableDrones(w http.ResponseWriter, r *http.Request) {
	var drones []drone.Drone
	drones, err := drone.AvailableDrones(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(drones)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
