package state

type DroneState string

const (
	IDLE       DroneState = "IDLE"
	LOADING    DroneState = "LOADING"
	LOADED     DroneState = "LOADED"
	DELIVERING DroneState = "DELIVERING"
	DELIVERED  DroneState = "DELIVERED"
	RETURNING  DroneState = "RETURNING"
)
