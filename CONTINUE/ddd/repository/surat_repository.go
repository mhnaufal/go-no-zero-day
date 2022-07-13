package repository

import (
	"context"
	"no-zero-day/CONTINUE/ddd/entity"
)

type SuratRepository interface {
	Insert(ctx context.Context, surat entity.Surat) (entity.Surat, error)
	FindById(ctx context.Context, id int32) (entity.Surat, error)
	FindAll(ctx context.Context) ([]entity.Surat, error)
}
