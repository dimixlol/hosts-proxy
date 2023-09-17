package usecases

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/models"
	"github.com/dimixlol/knowyourwebsite/ports"
	"github.com/dimixlol/knowyourwebsite/utils"
)

func NewPersistentSite(ctx context.Context, persister ports.Persister, host *models.Host, ip *models.IP) ports.URL {
	var err error

	err = persister.GetOrCreateWithTrack(host, &models.Host{Host: host.Host})
	if err != nil {
		panic(err)
	}

	err = persister.GetOrCreateWithTrack(ip, &models.IP{IP: ip.IP})
	if err != nil {
		panic(err)
	}

	urlModel := &models.URL{Host: host, IP: ip, Slug: utils.RandomStringWithLength(config.Configuration.SlugLength)}

	err = persister.CreateURL(urlModel)
	if err != nil {
		panic(err)
	}

	return urlModel
}
