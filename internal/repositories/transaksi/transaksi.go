package transaksi

import (
	"assesment-isi/config"
	"assesment-isi/internal/models"
	"fmt"
)

func (r *TransaksiRepository) GetNasabahByRekening(noRekening string) (result models.Nasabah, err error) {
	key := fmt.Sprintf(sfGetNasabahByRekening, noRekening)
	if _, err, _ = r.singleFlight.Do(key, func() (interface{}, error) {

		result, err = r.transaksi.GetNasabahByRekening(noRekening)
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"no_rekening": noRekening,
			})
			return nil, err
		}

		return nil, nil
	}); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
		})
		return result, err
	}

	return result, nil
}

func (r *TransaksiRepository) UpdateSaldoNasabah(nasabah models.Nasabah, nominal int64) (err error) {
	key := fmt.Sprintf(sfUpdateSaldoNasabah, nasabah)
	if _, err, _ = r.singleFlight.Do(key, func() (interface{}, error) {

		err = r.transaksi.UpdateSaldoNasabah(nasabah, nominal)
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

func (r *TransaksiRepository) CreateTransaksi(transaksi models.Transaksi) (err error) {
	key := fmt.Sprintf(sfCreateTransaksi, transaksi)
	if _, err, _ = r.singleFlight.Do(key, func() (interface{}, error) {

		err = r.transaksi.CreateTransaksi(transaksi)
		if err != nil {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"transaksi": transaksi,
			})
			return nil, err
		}

		return nil, nil
	}); err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"transaksi": transaksi,
		})
		return err
	}

	return nil
}
