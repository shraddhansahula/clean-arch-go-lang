package domain

type Vehicle struct{
	ID    uint64 `json:"id" gorm:"primary_key"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type VehicleInput struct{
	Brand  string `json:"brand" binding:"required"`
	Model string `json:"model" binding:"required"`
}
