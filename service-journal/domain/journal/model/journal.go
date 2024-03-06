package model

import "time"

type Journal struct {
	ID               int64     `json:"id" gorm:"primaryKey;column:id;type:bigserial"`
	TanggalTramsaksi time.Time `json:"tanggal_transaksi" gorm:"column:tanggal_transaksi;type:timestamp"`
	NoRekeningKredit int64     `json:"no_rekening_kredit" gorm:"column:no_rekening_kredit;type:integer;default:0"`
	NoRekeningDebit  int64     `json:"no_rekening_debit" gorm:"column:no_rekening_debit;type:integer;default:0"`
	NominalKredit    int64     `json:"nominal_kredit" gorm:"column:nominal_kredit;type:integer;default:0"`
	NominalDebit     int64     `json:"nominal_debit" gorm:"column:nominal_debit;type:integer;default:0"`
}

func (Journal) TableName() string {
	return "public.journal"
}

type CreateJournalRequest struct {
	TanggalTramsaksi time.Time `json:"tanggal_transaksi"`
	NoRekeningKredit int64     `json:"no_rekening_kredit"`
	NoRekeningDebit  int64     `json:"no_rekening_debit"`
	NominalKredit    int64     `json:"nominal_kredit"`
	NominalDebit     int64     `json:"nominal_debit"`
}
