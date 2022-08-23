package drone

import (
	"drone/medication"

	"gorm.io/gorm"
)

// type Drone struct {
// 	DroneID         int
// 	Status          string
// 	BatteryCapacity int
// 	Medications     []medication.Medication
// }

func GetDrones(db *gorm.DB) ([]Drone, error) {

	var drones []Drone

	if result := db.Find(&drones); result.Error != nil {
		return []Drone{}, result.Error
	}
	return drones, nil
}

func GetDrone(db *gorm.DB, id int) (Drone, error) {
	var drone Drone

	if result := db.First(&drone, id); result.Error != nil {
		return Drone{}, result.Error
	}
	return drone, nil
}

func RegisterDrone(db *gorm.DB, d Drone) (Drone, error) {

	if result := db.Create(&d); result.Error != nil {
		return Drone{}, result.Error
	}
	return d, nil
}

func (d *Drone) CheckingLoadedMedication(db *gorm.DB) ([]medication.Medication, error) {
	var medications []medication.Medication
	db.Where("drone_id = ?", d.ID).Find(&medications)
	return medications, nil
}

func (d *Drone) LoadMedication(medications []medication.Medication) error {
	// return d medication
	return nil
}
