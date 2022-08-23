package medication

type Medication struct {
	MedicationCode string `json:"medication_code" gorm:"primaryKey"`
	DroneID        int    `gorm:"foreignKey:DroneID"`
}

func (Medication) TableName() string {
	return `"drone"."medications"`
}
