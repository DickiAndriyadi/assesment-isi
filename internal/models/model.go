package models

import (
	"time"

	"gorm.io/gorm"
)

type Nasabah struct {
	gorm.Model
	Nama       string `json:"nama" gorm:"column:nama;not null;type:varchar(191)"`
	NIK        string `json:"nik" gorm:"column:nik;uniqueIndex:nik_type_unique;not null;type:varchar(191)"`
	NoHP       string `json:"no_hp" gorm:"column:no_hp;uniqueIndex:no_hp_type_unique;not null;type:varchar(191)"`
	NoRekening string `json:"no_rekening" gorm:"column:no_rekening;uniqueIndex;not null;type:varchar(191)"`
	Saldo      int64  `json:"saldo" gorm:"column:saldo;default:0"`

	Transaksi []*Transaksi `gorm:"foreignKey:NasabahID"`
}

func (Nasabah) TableName() string {
	return "nasabah"
}

type Transaksi struct {
	gorm.Model
	NasabahID      uint      `json:"nasabah_id" gorm:"column:nasabah_id;not null"`
	NoRekening     string    `json:"no_rekening" gorm:"column:no_rekening;not null"`
	Tipe           string    `json:"tipe" gorm:"column:tipe_transaksi;not null"`
	Nominal        int64     `json:"nominal" gorm:"column:nominal;not null"`
	WaktuTransaksi time.Time `json:"waktu_transaksi" gorm:"column:waktu_transaksi;not null"`

	// Belongs to
	Nasabah Nasabah `gorm:"foreignKey:NasabahID;references:ID"`
}

func (Transaksi) TableName() string {
	return "transaksi"
}
