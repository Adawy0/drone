package drone

import (
	Dronestate "drone/State"
	"drone/medication"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetDrones(db *gorm.DB) ([]Drone, error) {

	var drones []Drone

	if result := db.Find(&drones); result.Error != nil {
		return []Drone{}, result.Error
	}
	return drones, nil
}

func CheckBattery(db *gorm.DB, id int) (int, error) {
	var drone Drone
	if result := db.First(&drone, id); result.Error != nil {
		return 0, result.Error
	}
	return drone.BatteryCapacity, nil
}

func AvailableDrones(db *gorm.DB) ([]Drone, error) {
	var drones []Drone
	if result := db.Where("medication IS NULL OR medication = '{}'").Find(&drones); result.Error != nil {
		return []Drone{}, result.Error
	}
	return drones, nil
}

func RegisterDrone(db *gorm.DB, d Drone) (Drone, error) {
	if d.Weight > 500 {
		return Drone{}, errors.New("max drone weight is 500gr")
	}
	if result := db.Create(&d); result.Error != nil {
		return Drone{}, result.Error
	}
	return d, nil
}

func CheckingLoadedMedication(db *gorm.DB, id int) ([]medication.Medication, error) {
	d, err := getDrone(db, id)
	if err != nil {
		return []medication.Medication{}, err
	}
	var medications []medication.Medication
	db.Where("drone_id = ?", d.ID).Find(&medications)
	return medications, nil
}

func LoadMedication(db *gorm.DB, medications []medication.Medication, id int) error {
	d, err := getDrone(db, id)
	if err != nil {
		return err
	}
	if d.BatteryCapacity <= 25 {
		return errors.New(`can not loading medication for this drone, battery level is **below 25%**`)
	}
	d.State = string(Dronestate.LOADING)
	currentPayload := 0
	for _, medication := range d.Medications {
		currentPayload += medication.Weight
	}
	// assumptions slice is sorted
	for _, medication := range medications {
		if currentPayload+medication.Weight < int(d.Weight) {
			d.Medications = append(d.Medications, medication)
			// simulation real LOADING
			time.Sleep(10 * time.Second)
		} else {
			break
		}

	}
	d.State = string(Dronestate.LOADED)
	return nil
}

func getDrone(db *gorm.DB, id int) (Drone, error) {
	var drone Drone

	if result := db.First(&drone, id); result.Error != nil {
		return Drone{}, result.Error
	}
	return drone, nil
}

func (d *Drone) Fly() {
	d.checkBattery()
	d.checkPropeller()
	d.healthCheck()
}

func (d *Drone) checkBattery() {
	fmt.Println("[preparing] checking battery's status ... ")
}

func (d *Drone) checkPropeller() {
	fmt.Println("[preparing] checking propellers' status ... ")
}

func (d *Drone) healthCheck() {
	fmt.Println("[flying] on the air, everything is ok, auto balancing enabled ... ")
}
