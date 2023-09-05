package ports

import "github.com/dimixlol/knowyourwebsite/domains/persister/models"

type TrackedModel interface {
	IncrementTimes()
}

type URL interface {
	GetIP() string
	GetHost() string
	GetSlug() string
}

type Persister interface {
	NewPersistentSite(host *models.Host, ip *models.IP) URL
	GetURLBySlug(slug string) (URL, error)
}
