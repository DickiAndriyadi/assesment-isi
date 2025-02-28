package nasabah

import (
	"assesment-isi/config"
	"assesment-isi/internal/models"
	"fmt"
)

func (r *NasabahRepository) CheckExistingNasabah(nik string, noHP string) (result models.Nasabah, err error) {
	key := fmt.Sprintf(sfCheckExistingNasabah, nik, noHP)
	if _, err, _ = r.singleFlight.Do(key, func() (interface{}, error) {

		result, err = r.nasabah.CheckExistingNasabah(nik, noHP)
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nik":   nik,
				"no_hp": noHP,
			})
			return nil, err
		}

		return nil, nil
	}); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	return result, nil
}

func (r *NasabahRepository) CreateNasabah(nasabah models.Nasabah) (err error) {
	key := fmt.Sprintf(sfCreateNasabah, nasabah)
	if _, err, _ = r.singleFlight.Do(key, func() (interface{}, error) {

		err = r.nasabah.CreateNasabah(nasabah)
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nasabah": nasabah,
			})
			return nil, err
		}

		return nil, nil
	}); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nasabah": nasabah,
		})
		return err
	}

	return nil
}
