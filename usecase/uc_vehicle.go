package usecase

import (
	"context"
	"vehicles-service/domain"
	"vehicles-service/repository/sql"
)

type vehicleUsecase struct{
	vehicleRepo repo.VehicleRepository
}

func NewVehicleUsecase(repo repo.VehicleRepository) VehicleUsecase{
	return &vehicleUsecase{repo}
}

func (uc *vehicleUsecase) CreateNewVehicle(ctx context.Context, v *domain.Vehicle) error{
	uc.vehicleRepo.AddVehicle(ctx, v)
	return nil
}
func (uc *vehicleUsecase) GetVehicleByID(ctx context.Context, id uint64) (domain.Vehicle, error){
	v, _ := uc.vehicleRepo.SearchVehicleByID(ctx, id)
	return v, nil
}

func (uc *vehicleUsecase) ChangeVehicle(ctx context.Context, v *domain.Vehicle) error{
	return nil
}

