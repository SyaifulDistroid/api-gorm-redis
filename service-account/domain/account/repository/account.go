package store

import (
	"context"
	"service-account/domain/account/model"

	"gorm.io/gorm"
)

type AccountRepository interface {
	CreateRekeningRepository(ctx context.Context, rek *model.Account) (err error)
	GetByNIK(ctx context.Context, nik string) (rekening *model.Account, err error)
	GetByNoHandphone(ctx context.Context, hp string) (rekening *model.Account, err error)
	GetByNoRekening(ctx context.Context, norek int64) (rekening *model.Account, err error)
	UpdateSaldoRepository(ctx context.Context, rek *model.Account) (err error)
	CreateHistoryRepository(ctx context.Context, rek *model.History) (err error)
	GetHistoryRepository(ctx context.Context, norek int64) (rekening []*model.History, err error)
	GetAllHistoryRepository(ctx context.Context) (rekening []*model.History, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(Conn *gorm.DB) AccountRepository {
	return &accountRepository{
		db: Conn,
	}
}

func (acc *accountRepository) CreateRekeningRepository(ctx context.Context, rek *model.Account) (err error) {
	conn := acc.db.WithContext(ctx)
	return conn.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table(model.Account{}.TableName()).Create(&model.Account{
			Nama:        rek.Nama,
			NIK:         rek.NIK,
			PIN:         rek.PIN,
			NoHandphone: rek.NoHandphone,
			NoRekening:  rek.NoRekening,
			Saldo:       rek.Saldo,
		}).Error
		if err != nil {
			return
		}
		return
	})
}

func (acc *accountRepository) GetByNIK(ctx context.Context, nik string) (rekening *model.Account, err error) {
	query := acc.db.WithContext(ctx)
	query = query.Model(&model.Account{})
	query = query.Table(model.Account{}.TableName()).Where("nik = ?", nik)
	err = query.First(&rekening).Error
	if err != nil {
		return
	}

	return
}

func (acc *accountRepository) GetByNoHandphone(ctx context.Context, hp string) (rekening *model.Account, err error) {
	query := acc.db.WithContext(ctx)
	query = query.Model(&model.Account{})
	query = query.Table(model.Account{}.TableName()).Where("no_handphone = ?", hp)
	err = query.First(&rekening).Error
	if err != nil {
		return
	}

	return
}

func (acc *accountRepository) GetByNoRekening(ctx context.Context, norek int64) (rekening *model.Account, err error) {
	query := acc.db.WithContext(ctx)
	query = query.Model(&model.Account{})
	query = query.Table(model.Account{}.TableName()).Where("no_rekening = ?", norek)
	err = query.First(&rekening).Error
	if err != nil {
		return
	}

	return
}

func (acc *accountRepository) UpdateSaldoRepository(ctx context.Context, rek *model.Account) (err error) {
	conn := acc.db.WithContext(ctx)
	return conn.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table(model.Account{}.TableName()).Where("id = ?", rek.ID).Updates(&model.Account{
			Saldo: rek.Saldo,
		}).Error
		if err != nil {
			return
		}
		return
	})
}

func (acc *accountRepository) CreateHistoryRepository(ctx context.Context, rek *model.History) (err error) {
	conn := acc.db.WithContext(ctx)
	return conn.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table(model.History{}.TableName()).Create(&model.History{
			NoRekening:    rek.NoRekening,
			KodeTransaksi: rek.KodeTransaksi,
			Nominal:       rek.Nominal,
			Saldo:         rek.Saldo,
			CreatedAt:     rek.CreatedAt,
		}).Error

		if err != nil {
			return
		}
		return
	})
}

func (acc *accountRepository) GetHistoryRepository(ctx context.Context, norek int64) (rekening []*model.History, err error) {
	query := acc.db.WithContext(ctx)
	query = query.Model(&model.History{})
	query = query.Table(model.History{}.TableName()).Where("no_rekening = ?", norek).Order("created_at desc")
	err = query.Find(&rekening).Error
	if err != nil {
		return
	}

	return
}

func (acc *accountRepository) GetAllHistoryRepository(ctx context.Context) (rekening []*model.History, err error) {
	query := acc.db.WithContext(ctx)
	query = query.Model(&model.History{})
	query = query.Table(model.History{}.TableName())
	err = query.Find(&rekening).Error
	if err != nil {
		return
	}

	return
}
