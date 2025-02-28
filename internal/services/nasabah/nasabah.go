package nasabah

import (
	"assesment-isi/config"
	"assesment-isi/internal/constant"
	"assesment-isi/internal/models"
	"assesment-isi/util"
)

func (s *NasabahService) RegisterNasabah(nama, nik, noHP string) (result models.Nasabah, err error) {
	_, err = s.nasabah.CheckExistingNasabah(nik, noHP)
	if err != nil {
		if err.Error() != constant.NasabahNotFound.Error() {
			config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
				"nik":   nik,
				"no_hp": noHP,
			})
			return result, err
		}
	}

	newNasabah := models.Nasabah{
		Nama:       nama,
		NIK:        nik,
		NoHP:       noHP,
		NoRekening: util.GenerateNoRekening(),
		Saldo:      0,
	}

	err = s.nasabah.CreateNasabah(newNasabah)
	if err != nil {
		config.PrintErrorLog(err, config.GetErrorFileLine(), map[string]interface{}{
			"nik":   nik,
			"no_hp": noHP,
		})
		return result, err
	}

	return newNasabah, nil
}
