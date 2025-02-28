package transaksi

import (
	"assesment-isi/internal/models"
	"assesment-isi/internal/repositories/transaksi/postgre"

	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

type ITransaksiRepo interface {
	GetNasabahByRekening(noRekening string) (models.Nasabah, error)
	UpdateSaldoNasabah(nasabah models.Nasabah, nominal int64) (err error)
	CreateTransaksi(transaksi models.Transaksi) error
}

type TransaksiRepository struct {
	singleFlight singleflight.Group
	transaksi    postgre.ITransaksiPostgreRepository
}

func InitTransaksiRepository(db *gorm.DB) ITransaksiRepo {
	return &TransaksiRepository{
		singleFlight: singleflight.Group{},
		transaksi:    postgre.InitTransaksiPostgreRepository(db),
	}
}
