package drone

import "drone/medication"

type Drone struct {
	ID              int     `json:"-" gorm:"primaryKey"`
	SerialNumber    string  `json:"serial_number" gorm:"size:100"`
	Weight          float32 `json:"weight" gorm:"size:500"`
	State           string  `json:"state" gorm:"default:IDLE"`
	Model           string  `json:"model"`
	BatteryCapacity int     `json:"battery_capactiy" gorm:"default:100"`
	Medications     []medication.Medication
}

func (Drone) TableName() string {
	return `"drone"."drones"`
}
