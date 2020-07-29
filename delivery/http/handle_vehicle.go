package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vehicles-service/domain"
	"vehicles-service/usecase"
)


type vehicleHandler struct {
	vehicleUsecase usecase.VehicleUsecase
}

func NewVehicleHandler(r *gin.RouterGroup, us usecase.VehicleUsecase) {

	handler := &vehicleHandler{
		vehicleUsecase: us,
	}
	r.GET("/get/:id", handler.getById)
	r.POST("/add", handler.addVehicle) //curl --header "Content-Type: application/json" --request POST --data '{"brand":"Ferrari","model":"Spider"}' http://localhost:8080/v1/add
	//mux router
	//database migrations
}

func (vh *vehicleHandler) getById(ctx *gin.Context){
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 32)
	vehicles, _ := vh.vehicleUsecase.GetVehicleByID(ctx.Request.Context(), id)
	ctx.JSON(200, gin.H{"data": vehicles})
}

func (vh *vehicleHandler) addVehicle(ctx *gin.Context){
	var input domain.VehicleInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vehicle := domain.Vehicle{Brand: input.Brand, Model: input.Model}
	vh.vehicleUsecase.CreateNewVehicle(ctx, &vehicle)
	ctx.JSON(http.StatusOK, gin.H{"data": vehicle})
}

