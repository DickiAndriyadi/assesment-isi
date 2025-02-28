package routes

import (
	"assesment-isi/internal/controllers/nasabah"
	"assesment-isi/internal/controllers/transaksi"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, nasabahCtrl nasabah.INasabahController, transaksiCtrl transaksi.ITransaksiController) {
	api := e.Group("/api")
	api.POST("/daftar", nasabahCtrl.RegisterNasabah)
	api.POST("/tabung", transaksiCtrl.Tabung)
	api.POST("/tarik", transaksiCtrl.Tarik)
	api.GET("/saldo/:no_rekening", transaksiCtrl.GetSaldo)
}
