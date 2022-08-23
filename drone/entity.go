package drone

import (
	"drone/medication"
)

type Drone struct {
	ID              int    `json:"id" gorm:"primaryKey"`
	Status          string `json:"status" gorm:"default:IDLE"`
	BatteryCapacity int    `json:"battery_capactiy" gorm:"default:100"`
	Medications     []medication.Medication
}

func (Drone) TableName() string {
	return `"drone"."drones"`
}
