package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_delivery "vehicles-service/delivery/http"
	"vehicles-service/domain"
	_repo "vehicles-service/repository/sql"
	_usecase "vehicles-service/usecase"
)

func main(){
	r := gin.Default()
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	db.DropTable(&domain.Vehicle{})
	db.AutoMigrate(&domain.Vehicle{})

	db.Create(&domain.Vehicle{Brand: "Porsche", Model: "911"})
	db.Create(&domain.Vehicle{Brand: "Ford", Model: "Mustang GT"})
	db.Create(&domain.Vehicle{Brand: "BMW", Model: "X3"})
	db.Create(&domain.Vehicle{Brand: "Audi", Model: "Q3"})
	db.Create(&domain.Vehicle{Brand: "Volkswagen", Model: "Polo"})

	repo := _repo.NewSqlVehicleRepository(*db)
	uc := _usecase.NewVehicleUsecase(repo)

	api := r.Group("/v1")

	_delivery.NewVehicleHandler(api, uc)

	r.Run()

}

