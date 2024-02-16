package ports

import "github.com/dimixlol/hosts-proxy/models"

type (
	CacheManager interface {
		GetUrlBySlug(slug string) (URL, error)
	}

	Requester interface {
		MakeRequest() (string, error)
	}

	TrackedModel interface {
		IncrementTimes()
	}

	URL interface {
		GetIP() string
		GetHost() string
		GetSlug() string
	}

	Storage interface {
		GetOrCreateWithTrack(instance TrackedModel, model interface{}) error
		GetURLBySlug(slug string) (URL, error)
		CreateURL(url *models.URL) error
	}
)
