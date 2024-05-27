package model

import "time"

type Mutasi struct {
	ID               int64     `json:"id" gorm:"primaryKey;column:id;type:bigserial"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" gorm:"column:tanggal_transaksi;type:timestamp"`
	NoRekening       int64     `json:"no_rekening" gorm:"column:no_rekening;type:integer;default:0"`
	JenisTransaksi   string    `json:"jenis_transaksi" gorm:"column:jenis_transaksi;type:varchar(10)"`
	Nominal          int64     `json:"nominal" gorm:"column:nominal;type:integer;default:0"`
}

func (Mutasi) TableName() string {
	return "public.mutasi"
}

type CreateMutasiRequest struct {
	TanggalTramsaksi time.Time `json:"tanggal_transaksi"`
	NoRekening       int64     `json:"no_rekening"`
	JenisTransaksi   string    `json:"jenis_transaksi"`
	Nominal          int64     `json:"nominal"`
}
