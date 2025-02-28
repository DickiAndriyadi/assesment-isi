package postgre

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"assesment-isi/internal/models"
	"context"
	"strings"

	"github.com/afex/hystrix-go/hystrix"
	"gorm.io/gorm"
)

func (r *NasabahPostgreRepository) CheckExistingNasabah(nik, noHP string) (result models.Nasabah, err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {
		query := r.db.Where("nik = ? OR no_hp = ?", nik, noHP).First(&result)
		err = query.Error

		if err != nil {

			if err.Error() == gorm.ErrRecordNotFound.Error() {
				err = constant.NasabahNotFound
			}

			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nik":   nik,
				"no_hp": noHP,
			})
			return err
		}

		return nil
	}, nil); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	return result, nil
}

func (r *NasabahPostgreRepository) CreateNasabah(nasabah models.Nasabah) (err error) {
	if err := hystrix.DoC(context.Background(), constant.HystrixPostgre, func(ctx context.Context) error {

		err = r.db.Create(&nasabah).Error
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = constant.NasabahHasAlreadyExist
			}

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
