package main

import (
	"assesment-isi/config"
	"assesment-isi/routes"
	"log"

	nasabahController "assesment-isi/internal/controllers/nasabah"
	nasabahRepo "assesment-isi/internal/repositories/nasabah"
	nasabahService "assesment-isi/internal/services/nasabah"

	transaksiController "assesment-isi/internal/controllers/transaksi"
	transaksiRepo "assesment-isi/internal/repositories/transaksi"
	transaksiService "assesment-isi/internal/services/transaksi"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load environment & database
	config.ConnectDB()
	config.MigrateDB()

	// Argument parser untuk host & port
	configs := config.ParseArgs()

	// Inisialisasi repository, service, dan controller
	nRepository := nasabahRepo.InitNasabahRepository(config.DB)
	nService := nasabahService.InitNasabahService(nRepository)
	nController := nasabahController.InitNasabahController(nService)

	tRepository := transaksiRepo.InitTransaksiRepository(config.DB)
	tService := transaksiService.InitTransaksiService(tRepository)
	tController := transaksiController.InitTransaksiController(tService)

	// Inisialisasi Echo
	e := echo.New()

	// Routing
	routes.InitRoutes(e, nController, tController)

	// Jalankan server dengan argument parser
	serverAddr := configs.Host + ":" + configs.Port
	log.Println("Server berjalan di", serverAddr)
	e.Logger.Fatal(e.Start(serverAddr))
}
