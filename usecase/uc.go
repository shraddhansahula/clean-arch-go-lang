package usecase

import (
	"context"
	"vehicles-service/domain"
)

type VehicleUsecase interface{
	CreateNewVehicle(ctx context.Context, v *domain.Vehicle) error
	GetVehicleByID(ctx context.Context, id uint64) (domain.Vehicle, error)
	ChangeVehicle(ctx context.Context, v *domain.Vehicle) error
}

