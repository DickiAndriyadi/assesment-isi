package constant

import (
	"assesment-isi/internal/models"
	"net/http"
)

var (

	// 404 Not Found
	NasabahNotFound = models.WrapError(http.StatusNotFound, "Nasabah tidak ditemukan!")

	// 400 Bad Request
	NIKOrNoHPAlreadyExist  = models.WrapError(http.StatusBadRequest, "NIK atau No HP sudah digunakan!")
	NasabahHasAlreadyExist = models.WrapError(http.StatusBadRequest, "Nasabah sudah ada!")
	InvalidRequest         = models.WrapError(http.StatusBadRequest, "Invalid request!")
	ZeroNominal            = models.WrapError(http.StatusBadRequest, "Nominal harus lebih dari 0!")
	LowBalance             = models.WrapError(http.StatusBadRequest, "Saldo tidak cukup!")
)

type RestResponse struct {
	Remark string      `json:"remark"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}
