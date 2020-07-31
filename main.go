package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"

	//"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	_delivery "vehicles-service/delivery/http"
	"vehicles-service/domain"
	_repo "vehicles-service/repository/sql"
	_usecase "vehicles-service/usecase"
)

func main(){
	//r := gin.Default()
	user, _ := os.LookupEnv("POSTGRES_USER")
	port, _ := os.LookupEnv("POSTGRES_PORT")
	dbName, _ := os.LookupEnv("POSTGRES_DB")
	dbPassword, _ := os.LookupEnv("POSTGRES_PASSWORD")
	host, _ := os.LookupEnv("POSTGRES_HOST")
	postgresConname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", host, port, user, dbName, dbPassword)
	r := mux.NewRouter()
	db, err := gorm.Open("postgres", postgresConname)
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

	//api := r.Group("/v1")

	_delivery.NewVehicleHandler(r, uc)

	//r.Run()
	http.ListenAndServe(":8080", r)
}

//add db migrations
//add authentication headers
