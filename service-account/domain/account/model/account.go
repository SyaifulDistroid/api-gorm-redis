package model

import "time"

type Account struct {
	ID          int64  `json:"id" gorm:"primaryKey;column:id;type:bigserial"`
	Nama        string `json:"nama" gorm:"column:nama;type:varchar(200)"`
	NIK         string `json:"nik" gorm:"column:nik;type:varchar(200)"`
	PIN         string `json:"pin" gorm:"column:pin;type:varchar(200)"`
	NoHandphone string `json:"no_handphone" gorm:"column:no_handphone;type:varchar(200)"`
	NoRekening  int64  `json:"no_rekening" gorm:"column:no_rekening;type:integer;default:0"`
	Saldo       int64  `json:"saldo" gorm:"column:saldo;type:integer;default:0"`
}

func (Account) TableName() string {
	return "public.account"
}

type CreateAccountRequest struct {
	Nama        string `json:"nama"`
	NIK         string `json:"nik"`
	NoHandphone string `json:"no_handphone"`
	PIN         string `json:"pin"`
}

type SavingAccountRequest struct {
	NoRekening int64 `json:"no_rekening"`
	Nominal    int64 `json:"nominal"`
}

type WitdrawalAccountRequest struct {
	NoRekening int64 `json:"no_rekening"`
	Nominal    int64 `json:"nominal"`
}

type NasabahJSON struct {
	NoRekening int64 `json:"no_rekening"`
}

type SaldoJSON struct {
	NoRekening int64 `json:"no_rekening"`
	Saldo      int64 `json:"saldo"`
}

type History struct {
	ID            int64     `json:"id" gorm:"primaryKey;column:id;type:bigserial"`
	NoRekening    int64     `json:"no_rekening" gorm:"column:no_rekening;type:integer;default:0"`
	KodeTransaksi string    `json:"kode_transaksi" gorm:"column:kode_transaksi;type:varchar(200)"`
	Nominal       int64     `json:"nominal" gorm:"column:nominal;type:integer;default:0"`
	Saldo         int64     `json:"saldo" gorm:"column:saldo;type:integer;default:0"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:timestamp"`
}

func (History) TableName() string {
	return "public.history"
}

type HistoryJSON struct {
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       int64     `json:"nominal"`
	Saldo         int64     `json:"saldo"`
	Waktu         time.Time `json:"waktu"`
}

type TransferRequest struct {
	NoRekeningAsal   int64 `json:"no_rekening_asal"`
	NoRekeningTujuan int64 `json:"no_rekening_tujuan"`
	Nominal          int64 `json:"nominal"`
}

type Journal struct {
	TanggalTramsaksi time.Time `json:"tanggal_transaksi"`
	NoRekeningKredit int64     `json:"no_rekening_kredit"`
	NoRekeningDebit  int64     `json:"no_rekening_debit"`
	NominalKredit    int64     `json:"nominal_kredit"`
	NominalDebit     int64     `json:"nominal_debit"`
}

type GenerateToken struct {
	NoRekening int64  `json:"no_rekening"`
	PIN        string `json:"pin"`
}

type Token struct {
	Token string `json:"token"`
}
