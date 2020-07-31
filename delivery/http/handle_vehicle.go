package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"vehicles-service/domain"
	"vehicles-service/usecase"
)


type vehicleHandler struct {
	vehicleUsecase usecase.VehicleUsecase
}

func NewVehicleHandler(r *mux.Router, us usecase.VehicleUsecase) {

	handler := &vehicleHandler{
		vehicleUsecase: us,
	}
	r.HandleFunc("/get/{id}", handler.getById).Methods("GET")
	r.HandleFunc("/add", handler.addVehicle).Methods("POST")

	//gin code
	//r.GET("/get/:id", handler.getById)
	//r.POST("/add", handler.addVehicle) //curl --header "Content-Type: application/json" --request POST --data '{"brand":"Ferrari","model":"Spider"}' http://localhost:8080/add
}

func (vh *vehicleHandler) getById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 32)
	vehicles, _ := vh.vehicleUsecase.GetVehicleByID(r.Context(), id)
	payload, _ := json.Marshal(vehicles)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	fmt.Println(vehicles)
}

func (vh *vehicleHandler) addVehicle(w http.ResponseWriter, r *http.Request){
	var input domain.VehicleInput
	_ = json.NewDecoder(r.Body).Decode(&input)
	fmt.Println(input)
	vehicle := domain.Vehicle{Brand: input.Brand, Model: input.Model}
	vh.vehicleUsecase.CreateNewVehicle(r.Context(), &vehicle)
	fmt.Println(vehicle)
}

// gin code
//func (vh *vehicleHandler) getById(ctx *gin.Context){
//	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
//	vehicles, _ := vh.vehicleUsecase.GetVehicleByID(ctx.Request.Context(), id)
//	ctx.JSON(200, gin.H{"data": vehicles})
//}
//
//func (vh *vehicleHandler) addVehicle(ctx *gin.Context){
//	var input domain.VehicleInput
//	if err := ctx.ShouldBindJSON(&input); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	vehicle := domain.Vehicle{Brand: input.Brand, Model: input.Model}
//	vh.vehicleUsecase.CreateNewVehicle(ctx, &vehicle)
//	ctx.JSON(http.StatusOK, gin.H{"data": vehicle})
//}

