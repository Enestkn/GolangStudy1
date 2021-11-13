package config

type Form struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`         // Personel Adı
	SurName     string `json:"sur_name"`     // Soyadı
	PhoneNumber string `json:"phone_number"` // Telefonu
}

var Personels []Form

func Init() {
	var personel Form
	personel.Name = "Ibrahim"
	personel.SurName = "COBANI"
	personel.PhoneNumber = "0532 540 1194"
	Personels = append(Personels, personel)

	personel = Form{
		Name:        "Enes",
		SurName:     "Tekne",
		PhoneNumber: "0532 653 1876",
	}
	Personels = append(Personels, personel)

	Personels = append(Personels, Form{
		Name:        "Ali Enes",
		SurName:     "Yaman",
		PhoneNumber: "0532 444 2233",
	})
}
