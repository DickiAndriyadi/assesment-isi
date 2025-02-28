package transaksi

import (
	"assesment-isi/internal/models"
	"assesment-isi/internal/repositories/transaksi"
)

type ITransaksiService interface {
	Tabung(noRekening string, nominal int64) (models.Nasabah, error)
	Tarik(noRekening string, nominal int64) (models.Nasabah, error)
	GetSaldo(noRekening string) (int64, error)
}

type TransaksiService struct {
	transaksi transaksi.ITransaksiRepo
}

func InitTransaksiService(transaksi transaksi.ITransaksiRepo) ITransaksiService {
	return &TransaksiService{
		transaksi: transaksi,
	}
}
