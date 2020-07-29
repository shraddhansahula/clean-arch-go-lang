package repo

import (
	"context"
	//"database/sql"
	"github.com/jinzhu/gorm"
	"vehicles-service/domain"
)
type sqlVehicleRepository struct{
	Conn gorm.DB
}

func NewSqlVehicleRepository(Conn gorm.DB) VehicleRepository {
	return &sqlVehicleRepository{Conn}
}

func (repo *sqlVehicleRepository) AddVehicle(ctx context.Context, v *domain.Vehicle) error {
	repo.Conn.Create(v)
	return nil
}

func (repo *sqlVehicleRepository) SearchVehicleByID(ctx context.Context, id uint64) (domain.Vehicle, error){
	var v domain.Vehicle
	repo.Conn.Where("ID = ?", id).First(&v)
	return v, nil
}

func (repo *sqlVehicleRepository) UpdateVehicle(ctx context.Context, v *domain.Vehicle) error{
	return nil
}