package nasabah

import (
	"assesment-isi/internal/models"
	"assesment-isi/internal/repositories/nasabah/postgre"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type INasabahRepo interface {
	CheckExistingNasabah(nik, noHP string) (models.Nasabah, error)
	CreateNasabah(nasabah models.Nasabah) error
}

type NasabahRepository struct {
	singleFlight singleflight.Group
	nasabah      postgre.INasabahPostgreRepository
}

func InitNasabahRepository(db *gorm.DB) INasabahRepo {
	return &NasabahRepository{
		singleFlight: singleflight.Group{},
		nasabah:      postgre.InitNasabahPostgreRepository(db),
	}
}
