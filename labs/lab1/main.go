package main

import "fmt"

func main() {
	type Form struct {
		Name        string `json:"name"`         // Personel Adı
		SurName     string `json:"sur_name"`     // Soyadı
		PhoneNumber string `json:"phone_number"` // Telefonu
	}

	var personels []Form
	var personel Form
	personel.Name = "Ibrahim"
	personel.SurName = "COBANI"
	personel.PhoneNumber = "0532 540 1194"
	personels = append(personels, personel)

	personel = Form{
		Name:        "Enes",
		SurName:     "Tekne",
		PhoneNumber: "0532 653 1876",
	}
	personels = append(personels, personel)

	personels = append(personels, Form{
		Name:        "Ali Enes",
		SurName:     "Yaman",
		PhoneNumber: "0532 444 2233",
	})

	fmt.Println("Merhaba Dünya", personels)
	fmt.Printf("%+v\n", personels)

	for i, form := range personels {
		fmt.Printf("Listemizin %d sırasında yer alan  %s %s isimli arkadaşımızın telefon numarası %s'dir\n",
			i+1, form.Name, form.SurName, form.PhoneNumber)
		body := fmt.Sprintf("Listemizin %d sırasında yer alan  %s %s isimli arkadaşımızın telefon numarası %s'dir",
			i+1, form.Name, form.SurName, form.PhoneNumber)

		fmt.Println("*****************")
		fmt.Println(body)
		fmt.Println("*****************")
	}
}
