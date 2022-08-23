package main

import (
	"drone/medication"

	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220808232129(txn *gorm.DB) {
	type Drone struct {
		ID              int    `json:"id" gorm:"primaryKey"`
		Status          string `json:"status" gorm:"default:IDLE"`
		BatteryCapacity int    `json:"battery_capactiy" gorm:"default:100"`
		Medications     []medication.Medication
	}
	txn.AutoMigrate(&Drone{})

}

// Down is executed when this migration is rolled back
func Down_20220808232129(txn *gorm.DB) {
	txn.Migrator().DropTable("drone")

}
