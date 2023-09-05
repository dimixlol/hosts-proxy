package database

import (
	"context"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister/models"
	"github.com/dimixlol/knowyourwebsite/domains/persister/ports"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/dimixlol/knowyourwebsite/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormPersister struct {
	ctx context.Context
	DB  *gorm.DB
}

func (g *gormPersister) getOrCreateWithTrack(instance ports.TrackedModel, model interface{}) error {
	err := g.DB.FirstOrCreate(instance, model).Error
	if err != nil {
		return err
	}
	instance.IncrementTimes()
	g.DB.Save(instance)
	return nil
}

func (g *gormPersister) NewPersistentSite(host *models.Host, ip *models.IP) ports.URL {
	var err error
	err = g.getOrCreateWithTrack(host, &models.Host{Host: host.Host})
	if err != nil {
		panic(err)
	}
	err = g.getOrCreateWithTrack(ip, &models.IP{IP: ip.IP})
	if err != nil {
		panic(err)
	}

	urlModel := &models.URL{Host: host, IP: ip, Slug: utils.RandomStringWithLength(config.Configuration.SlugLength)}
	err = g.DB.Create(urlModel).Error
	if err != nil {
		panic(err)
	}
	return urlModel
}

func (g *gormPersister) GetURLBySlug(slug string) (ports.URL, error) {
	url := &models.URL{}
	res := g.DB.Preload(clause.Associations).Find(url, "slug = ?", slug)
	fmt.Println("ERROR is", res.Error)
	fmt.Println("URL is", url)
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
