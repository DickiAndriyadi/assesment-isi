package postgre

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"assesment-isi/internal/models"
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"gorm.io/gorm"
)

func (r *TransaksiPostgreRepository) GetNasabahByRekening(noRekening string) (result models.Nasabah, err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {
		if err := r.db.Where("no_rekening = ?", noRekening).First(&result).Error; err != nil {

			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = constant.NasabahNotFound
			}

			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"no_rekening": noRekening,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
		})
		return result, err
	}

	return result, nil
}

func (j *TransaksiPostgreRepository) UpdateSaldoNasabah(nasabah models.Nasabah, nominal int64) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		nasabah.Saldo += nominal
		if err := j.db.Where("id = ?", nasabah.ID).Updates(&nasabah).Error; err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nasabah": nasabah,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nasabah": nasabah,
		})
		return err
	}

	return nil
}

func (r *TransaksiPostgreRepository) CreateTransaksi(transaksi models.Transaksi) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		err = r.db.Create(&transaksi).Error
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"transaksi": transaksi,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"transaksi": transaksi,
		})
		return err
	}

	return nil
}
