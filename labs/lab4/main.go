package main

import (
	"errors"
	"fmt"
	"github.com/enestkn/golangstudy1/labs/lab4/config"
	"github.com/enestkn/golangstudy1/labs/lab4/models"
	"gorm.io/gorm"
)

func main() {
	config.Init()
	if err := config.DB.AutoMigrate(&models.Personal{}); err != nil {
		fmt.Println(err)
	}

	//InsertRecordSample()

	foundPersonal := FindPersonalByEmail("enestekne0@gmail.com")
	fmt.Printf("Email %+v\n", foundPersonal)

	foundPersonal = FindPersonalByID(16)
	fmt.Printf("ID %+v\n", foundPersonal)

	personals := PersonalList()
	for i, personal := range personals {
		fmt.Printf("%d (%d) - %s %s %s %s\n",
			i, personal.ID,
			personal.Name, personal.SurName, personal.Email, personal.Phone)

		if personal.Email == "ibrahim@imaconsult.com" {
			personal.Name = "ICIBEY"
			UpdatePersonal(personal)
		}
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	personals = PersonalList()
	for i, personal := range personals {
		fmt.Printf("%d (%d) - %s %s %s %s\n",
			i, personal.ID,
			personal.Name, personal.SurName, personal.Email, personal.Phone)
	}

}

// Kayıt ekleme ile ilgili örnek
func InsertRecordSample() {
	var personal models.Personal
	var personals []models.Personal

	config.DB.Where("id <> -1").Delete(&models.Personal{})
	personal.Name = "Ibrahim"
	personal.Email = "ibrahim@imaconsult.com"
	personal.SurName = "COBANI"
	personal.Phone = "0532 540 1194"
	personals = append(personals, personal)

	personal = models.Personal{
		Name:    "Enes",
		Email:   "enestekne0@gmail.com",
		SurName: "TEKNE",
		Phone:   "0532 653 1876",
	}
	personals = append(personals, personal)

	personals = append(personals, models.Personal{
		Name:    "Ali Enes",
		SurName: "YAMAN",
		Email:   "alienes0@gmail.com",
		Phone:   "0532 444 2233",
	})

	config.DB.Create(&personals)
}

func FindPersonalByEmail(email string) models.Personal {
	var personal models.Personal

	config.DB.
		Where("email = ?", email).
		First(&personal)

	return personal
}
func FindPersonalByID(id uint64) models.Personal {
	var personal models.Personal
	config.DB.First(&personal, id)
	return personal
}

func PersonalList() []models.Personal {
	var personals []models.Personal
	config.DB.Find(&personals)
	return personals
}

func UpdatePersonal(personal models.Personal) {
	var data models.Personal
	if err := config.DB.First(&data, personal.ID).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("Kaydı bulamadık.")
		return
	}
	config.DB.Save(&personal)
}
