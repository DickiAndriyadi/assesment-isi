package postgre

import (
	"assesment-isi/internal/models"

	"gorm.io/gorm"
)

type ITransaksiPostgreRepository interface {
	GetNasabahByRekening(noRekening string) (models.Nasabah, error)
	UpdateSaldoNasabah(nasabah models.Nasabah, nominal int64) (err error)
	CreateTransaksi(transaksi models.Transaksi) error
}

type TransaksiPostgreRepository struct {
	db *gorm.DB
}

func InitTransaksiPostgreRepository(db *gorm.DB) ITransaksiPostgreRepository {
	return &TransaksiPostgreRepository{
		db: db,
	}
}
