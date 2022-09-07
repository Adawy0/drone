package medication

import (
	"gorm.io/gorm"
)

func RegisterMedication(db *gorm.DB, m Medication) (Medication, error) {
	//- name (allowed only letters, numbers, ‘-‘, ‘_’);
	//- code (allowed only upper case letters, underscore and numbers);
	if result := db.Create(&m); result.Error != nil {
		return Medication{}, result.Error
	}
	return m, nil
}
