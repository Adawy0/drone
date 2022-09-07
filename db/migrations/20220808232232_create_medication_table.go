package main

import (
	"gorm.io/gorm"
)

// Up is executed when this migration is applied
func Up_20220808232232(txn *gorm.DB) {
	type Medication struct {
		Name    string `json:"name"`
		Code    string `json:"code" gorm:"primaryKey"`
		Weight  int    `json:"weight"`
		DroneID int    `gorm:"foreignKey:DroneID"`
		Image   []byte `json:"image"`
	}
	txn.AutoMigrate(&Medication{})

}

// Down is executed when this migration is rolled back
func Down_20220808232232(txn *gorm.DB) {
	txn.Migrator().DropTable("medication")

}
