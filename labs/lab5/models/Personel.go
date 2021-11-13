package models

import "time"

/*
	Bu tablo personel kayıtlarını tutmak için kullanılacak.
*/

type Personal struct {
	ID        uint64 `json:"id" gorm:"primaryKey"` // ID
	Name      string `json:"name"`                 // Personelin ismi
	SurName   string `json:"sur_name"`             // Soyismi
	Email     string `json:"email"`                // Email adresi
	Phone     string `json:"phone"`                // Telefonu
	CreatedAt time.Time
	UpdatedAt time.Time
}
