package main

import (
	"fmt"
	"log"

	"github.com/SetyaK/BL-Onboarding3-Go-package"
	"github.com/SetyaK/BL-Onboarding3-Go-package/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	sess, err := database.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	m := database.Migration{Session: sess}
	_, err = m.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	pRepo := ministore.ProductRepository{Session: sess}

	pRepo.Add("sample product", "the sample product of my mini store", 1)
	_, c, err := pRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

}
