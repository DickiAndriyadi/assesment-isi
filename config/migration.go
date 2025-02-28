package config

import (
	"assesment-isi/internal/models"
	"log"
)

func MigrateDB() {
	err := DB.AutoMigrate(&models.Nasabah{}, &models.Transaksi{})
	if err != nil {
		log.Fatalf("Database migration failed: %s", err)
	}

	DB.Exec("ALTER TABLE transaksi ALTER COLUMN no_rekening TYPE VARCHAR(50);")

	log.Println("Database migrated successfully!")
}
