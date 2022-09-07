package medication

type Medication struct {
	Name    string `json:"name"`
	Code    string `json:"code" gorm:"primaryKey"`
	Weight  int    `json:"weight"`
	Image   []byte `json:"image"`
	DroneID int    `gorm:"foreignKey:DroneID"`
}

func (Medication) TableName() string {
	return `"drone"."medications"`
}
