package server

import (
	"drone/db"
	"drone/drone"
	"drone/medication"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func setup() {
	fmt.Println("load fixtures")
	var err error
	conn, err := db.Init()
	// begin a transaction
	dbClient = conn.Begin()
	db.LoadFixtures(dbClient)
	if err != nil {
		log.Println("Couldn't connect to database", err.Error())
		log.Fatal(err)
	}
}

func shutdown() {
	fmt.Println("rollback database ...")
	// rollback the transaction
	dbClient.Rollback()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGETDrones(t *testing.T) {
	t.Run("Test get drones successfully", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/drone/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		h := handler{
			DB: dbClient,
		}
		handler := http.HandlerFunc(h.GetDrones)

		handler.ServeHTTP(rr, req)
		got := rr.Body.String()
		want := "[{\"id\":1,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null},{\"id\":2,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null},{\"id\":3,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null},{\"id\":4,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null},{\"id\":5,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null},{\"id\":6,\"status\":\"IDLE\",\"battery_capactiy\":100,\"Medications\":null}]"

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})
}

func TestRegisterDrone(t *testing.T) {
	t.Run("test register drone successfully", func(t *testing.T) {
		payload := `{"Status": "IDLE", "BatteryCapacity": 100}`
		req, err := http.NewRequest("POST", "/api/drone", strings.NewReader(payload))
		if err != nil {
			t.Fatal(err)
		}
		h := handler{
			DB: dbClient,
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(h.RegisterDrone)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
		var registerDrone drone.Drone
		err = json.Unmarshal([]byte(rr.Body.String()), &registerDrone)
		if err != nil {
			t.Errorf("error while parsing response")
		}
		expected := drone.Drone{
			ID:              registerDrone.ID,
			Status:          "IDLE",
			BatteryCapacity: 100,
		}

		if !reflect.DeepEqual(registerDrone, expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestCheckingLoadedMedication(t *testing.T) {
	t.Run("test checking loaded medication for given drone without pass drone id", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/drone//check-medication", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		h := handler{
			DB: dbClient,
		}
		handler := http.HandlerFunc(h.CheckingLoadedMedication)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := `{"error":"Couldn't find id in request URL"}`

		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("test checking loaded medication for given drone", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/drone/3/check-medication", nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{
			"id": "5",
		})
		rr := httptest.NewRecorder()
		h := handler{
			DB: dbClient,
		}
		handler := http.HandlerFunc(h.CheckingLoadedMedication)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}

		var medications []medication.Medication
		err = json.Unmarshal([]byte(rr.Body.String()), &medications)
		if err != nil {
			t.Errorf("error while parsing response")
		}
		expected := []medication.Medication{
			{
				MedicationCode: "M1",
				DroneID:        5,
			},
			{
				MedicationCode: "M2",
				DroneID:        5,
			},
			{
				MedicationCode: "M3",
				DroneID:        5,
			},
		}

		if !reflect.DeepEqual(medications, expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}

	})

	t.Run("test checking loaded medication for given drone not exist", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/drone/1/check-medication", nil)
		if err != nil {
			t.Fatal(err)
		}
		req = mux.SetURLVars(req, map[string]string{
			"id": "1000",
		})
		rr := httptest.NewRecorder()
		h := handler{
			DB: dbClient,
		}
		handler := http.HandlerFunc(h.CheckingLoadedMedication)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})
}

func TestLoadedMedication(t *testing.T) {
	t.Run("test load medication for given drone", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/api/drone/1/load-medication", nil)
		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req, map[string]string{
			"id": "1",
		})
		rr := httptest.NewRecorder()
		h := handler{
			DB: dbClient,
		}
		handler := http.HandlerFunc(h.LoadMedication)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// expected := `{"error":"Couldn't find id in request URL"}`

		// if rr.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v",
		// 		rr.Body.String(), expected)
		// }
	})

}
