package store

import (
	"context"
	"gorm.io/gorm"
	"service-journal/domain/journal/model"
)

type JournalRepository interface {
	CreateJournalRepository(ctx context.Context, req *model.Journal) (err error)
}

type journalRepository struct {
	db *gorm.DB
}

func NewJournalRepository(Conn *gorm.DB) JournalRepository {
	return &journalRepository{
		db: Conn,
	}
}

func (journal *journalRepository) CreateJournalRepository(ctx context.Context, req *model.Journal) (err error) {
	conn := journal.db.WithContext(ctx)
	return conn.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table(model.Journal{}.TableName()).Create(&model.Journal{
			TanggalTramsaksi: req.TanggalTramsaksi,
			NoRekeningKredit: req.NoRekeningKredit,
			NoRekeningDebit:  req.NoRekeningDebit,
			NominalKredit:    req.NominalKredit,
			NominalDebit:     req.NominalDebit,
		}).Error
		if err != nil {
			return
		}
		return
	})
}
