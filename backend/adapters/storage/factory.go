package storage

import (
	"context"
	"github.com/dimixlol/hosts-proxy/adapters/storage/database/gorm"
	"github.com/dimixlol/hosts-proxy/ports"
)

func NewDatabaseStorage(ctx context.Context) ports.Storage {
	return gorm.New(ctx)
}
