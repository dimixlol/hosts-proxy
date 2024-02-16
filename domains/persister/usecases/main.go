package usecases

import (
	"context"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/models"
	"github.com/dimixlol/hosts-proxy/ports"
	"github.com/dimixlol/hosts-proxy/utils"
)

func NewPersistentSite(ctx context.Context, persister ports.Storage, host *models.Host, ip *models.IP) ports.URL {
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
