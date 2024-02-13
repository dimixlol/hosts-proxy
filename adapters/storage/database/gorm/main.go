package gorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/dimixlol/hosts-proxy/models"
	"github.com/dimixlol/hosts-proxy/ports"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Connector struct {
	ctx context.Context
	DB  *gorm.DB
}

func (g *Connector) GetOrCreateWithTrack(instance ports.TrackedModel, model interface{}) error {
	err := g.DB.FirstOrCreate(instance, model).Error
	if err != nil {
		return err
	}
	instance.IncrementTimes()
	g.DB.Save(instance)
	return nil
}

func (g *Connector) CreateURL(url *models.URL) error {
	res := g.DB.First(&url, "host_id = ? AND ip_id = ?", url.Host.ID, url.IP.ID)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return g.DB.Create(url).Error
	}

	return res.Error
}

func (g *Connector) GetURLBySlug(slug string) (ports.URL, error) {
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

func New(ctx context.Context) *Connector {
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

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err = db.AutoMigrate(&models.Host{}, &models.IP{}, &models.URL{})
	if err != nil {
		panic(err)
	}

	return &Connector{
		ctx: ctx,
		DB:  db,
	}
}
