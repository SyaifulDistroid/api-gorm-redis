package store

import (
	"context"
	"service-mutasi/domain/mutasi/model"

	"gorm.io/gorm"
)

type MutasiRepository interface {
	CreateMutasiRepository(ctx context.Context, req *model.Mutasi) (err error)
}

type mutasiRepository struct {
	db *gorm.DB
}

func NewMutasiRepository(Conn *gorm.DB) MutasiRepository {
	return &mutasiRepository{
		db: Conn,
	}
}

func (mutasi *mutasiRepository) CreateMutasiRepository(ctx context.Context, req *model.Mutasi) (err error) {
	conn := mutasi.db.WithContext(ctx)
	return conn.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Table(model.Mutasi{}.TableName()).Create(&model.Mutasi{
			TanggalTransaksi: req.TanggalTransaksi,
			NoRekening:       req.NoRekening,
			JenisTransaksi:   req.JenisTransaksi,
			Nominal:          req.NoRekening,
		}).Error
		if err != nil {
			return
		}
		return
	})
}
