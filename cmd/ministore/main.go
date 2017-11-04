package main

import (
	"fmt"

	"github.com/SetyaK/BL-Onboarding3-Go-package"
	"github.com/SetyaK/BL-Onboarding3-Go-package/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	sess, error := database.NewSession()
	if error != nil {
		fmt.Println(error)
	} else {
		pRepo := ministore.ProductRepository{Session: sess}
		_, c := pRepo.GetAll()
		fmt.Println(c)
	}
}
