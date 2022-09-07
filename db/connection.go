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

var FixturesDrones []drone.Drone

func LoadFixtures(db *gorm.DB) error {
	FixturesDrones = []drone.Drone{
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{
			{Code: "Q1"},
			{Code: "Q1"},
		}},
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{
			{Code: "M1", Weight: 350},
			{Code: "M2", Weight: 100},
			{Code: "M3", Weight: 400},
		}},
		{State: "IDLE", BatteryCapacity: 100, Medications: []medication.Medication{}},
	}

	if result := db.Create(&FixturesDrones); result.Error != nil {
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
