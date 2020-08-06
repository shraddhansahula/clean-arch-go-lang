package repo

import (
	"context"
	"fmt"
	"sync"

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

func (repo *sqlVehicleRepository) search(id uint64, group *sync.WaitGroup, a chan domain.Vehicle) {
	defer group.Done()
	fmt.Println("in function for ", id)
	var v domain.Vehicle
	repo.Conn.Where("ID = ?", id).Find(&v)
	fmt.Println("done function for done ", id)
	a <- v
}


func(repo *sqlVehicleRepository) GetList(ctx context.Context, ids []int, mode int) []domain.Vehicle{
	var v []domain.Vehicle
	if mode == 1 {
		repo.Conn.Where(ids).Find(&v)
	} else {
		var wg sync.WaitGroup

		a := make(chan domain.Vehicle)
		for _, value := range ids{
			wg.Add(1)
			fmt.Println("id", value)
			go repo.search(uint64(value), &wg, a)
		}
		for range ids{
			v = append(v, <-a)
		}
		wg.Wait()
	}
	return v
}