package nasabah

import (
	"assesment-isi/internal/models"
	"assesment-isi/internal/repositories/nasabah"
)

type INasabahService interface {
	RegisterNasabah(nama, nik, noHP string) (result models.Nasabah, err error)
}

type NasabahService struct {
	nasabah nasabah.INasabahRepo
}

func InitNasabahService(nasabah nasabah.INasabahRepo) INasabahService {
	return &NasabahService{
		nasabah: nasabah,
	}
}
