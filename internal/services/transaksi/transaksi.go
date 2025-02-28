package transaksi

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"assesment-isi/internal/models"
	"time"
)

func (s *TransaksiService) Tabung(noRekening string, nominal int64) (result models.Nasabah, err error) {
	if nominal <= int64(0) {
		err = constant.ZeroNominal
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	nasabah, err := s.transaksi.GetNasabahByRekening(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	err = s.transaksi.UpdateSaldoNasabah(nasabah, nominal)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	transaksi := models.Transaksi{
		NasabahID:      nasabah.ID,
		NoRekening:     noRekening,
		Tipe:           "tabung",
		Nominal:        nominal,
		WaktuTransaksi: time.Now(),
	}
	err = s.transaksi.CreateTransaksi(transaksi)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	result, err = s.transaksi.GetNasabahByRekening(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	return result, nil
}

func (s *TransaksiService) Tarik(noRekening string, nominal int64) (result models.Nasabah, err error) {
	if nominal <= 0 {
		err = constant.ZeroNominal
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	nasabah, err := s.transaksi.GetNasabahByRekening(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	if nasabah.Saldo < nominal {
		err = constant.LowBalance
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	err = s.transaksi.UpdateSaldoNasabah(nasabah, -nominal)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	transaksi := models.Transaksi{
		NasabahID:      nasabah.ID,
		NoRekening:     noRekening,
		Tipe:           "tarik",
		Nominal:        -nominal,
		WaktuTransaksi: time.Now(),
	}
	err = s.transaksi.CreateTransaksi(transaksi)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	result, err = s.transaksi.GetNasabahByRekening(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
			"nominal":     nominal,
		})
		return result, err
	}

	return result, nil
}

func (s *TransaksiService) GetSaldo(noRekening string) (int64, error) {
	nasabah, err := s.transaksi.GetNasabahByRekening(noRekening)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"no_rekening": noRekening,
		})
		return 0, err
	}
	return nasabah.Saldo, nil
}
