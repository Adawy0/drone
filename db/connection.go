package db

import (
	"drone/drone"
	"drone/medication"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=drone port=5432 sslmode=disable TimeZone=Africa/Cairo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func LoadFixtures(db *gorm.DB) error {
	fixturesDrones := []drone.Drone{
		{ID: 1, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{ID: 2, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{
			{MedicationCode: "Q1", DroneID: 2},
			{MedicationCode: "Q1", DroneID: 2},
		}},
		{ID: 3, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{ID: 4, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{ID: 5, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{
			{MedicationCode: "M1", DroneID: 5},
			{MedicationCode: "M2", DroneID: 5},
			{MedicationCode: "M3", DroneID: 5},
		}},
		{ID: 6, Status: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
	}

	if result := db.Create(&fixturesDrones); result.Error != nil {
		return result.Error
	}

	// fixturesMedication := []medication.Medication{
	// 	{MedicationCode: "M2", DroneID: 7},
	// }
	// if result := db.Create(&fixturesMedication); result.Error != nil {
	// 	return result.Error
	// }
	fmt.Println("Fixtures loaded ...")
	return nil
}
