package repo

import (
	"context"
	"vehicles-service/domain"
)

type VehicleRepository interface{
	AddVehicle(ctx context.Context, v *domain.Vehicle) error
	SearchVehicleByID(ctx context.Context, id uint64) (domain.Vehicle, error)
	UpdateVehicle(ctx context.Context, v *domain.Vehicle) error
	GetList(ctx context.Context, ids []int, mode int) []domain.Vehicle
}

