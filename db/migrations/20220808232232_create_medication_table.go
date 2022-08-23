package main

import (
	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220808232232(txn *gorm.DB) {
	type Medication struct {
		MedicationCode string `json:"medication_code" gorm:"primaryKey"`
		DroneID        int    `gorm:"foreignKey:DroneID"`
	}
	txn.AutoMigrate(&Medication{})

}

// Down is executed when this migration is rolled back
func Down_20220808232232(txn *gorm.DB) {
	txn.Migrator().DropTable("medication")

}
