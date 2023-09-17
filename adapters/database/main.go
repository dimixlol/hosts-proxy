package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/dimixlol/knowyourwebsite/models"
	"github.com/dimixlol/knowyourwebsite/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormPersister struct {
	ctx context.Context
	DB  *gorm.DB
}

func (g *gormPersister) GetOrCreateWithTrack(instance ports.TrackedModel, model interface{}) error {
	err := g.DB.FirstOrCreate(instance, model).Error
	if err != nil {
		return err
	}
	instance.IncrementTimes()
	g.DB.Save(instance)
	return nil
}

func (g *gormPersister) CreateURL(url *models.URL) error {
	return g.DB.Create(url).Error
}

func (g *gormPersister) GetURLBySlug(slug string) (ports.URL, error) {
	url := &models.URL{}
	res := g.DB.Preload(clause.Associations).First(url, "slug = ?", slug)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}

	return url, nil
}

func NewDatabasePersister(ctx context.Context) ports.Persister {
	logger := logging.GetLogger(ctx)
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
				config.Configuration.DB.Host,
				config.Configuration.DB.User,
				config.Configuration.DB.Password,
				config.Configuration.DB.Name,
				config.Configuration.DB.Port,
				config.Configuration.DB.SSLMode,
				config.Configuration.DB.Timezone,
			),
		),
		&gorm.Config{Logger: logger},
	)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Host{}, &models.IP{}, &models.URL{})
	if err != nil {
		panic(err)
	}

	return &gormPersister{
		ctx: ctx,
		DB:  db,
	}
}
