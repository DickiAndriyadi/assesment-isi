package postgre

import (
	"assesment-isi/internal/models"

	"gorm.io/gorm"
)

type INasabahPostgreRepository interface {
	CheckExistingNasabah(nik, noHP string) (models.Nasabah, error)
	CreateNasabah(nasabah models.Nasabah) error
}

type NasabahPostgreRepository struct {
	db *gorm.DB
}

func InitNasabahPostgreRepository(db *gorm.DB) INasabahPostgreRepository {
	return &NasabahPostgreRepository{
		db: db,
	}
}
