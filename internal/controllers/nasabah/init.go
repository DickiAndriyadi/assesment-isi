package nasabah

import (
	"assesment-isi/internal/services/nasabah"

	"github.com/labstack/echo/v4"
)

type INasabahController interface {
	RegisterNasabah(c echo.Context) (err error)
}

type NasabahController struct {
	nasabah nasabah.INasabahService
}

func InitNasabahController(nasabah nasabah.INasabahService) INasabahController {
	return &NasabahController{
		nasabah: nasabah,
	}
}
