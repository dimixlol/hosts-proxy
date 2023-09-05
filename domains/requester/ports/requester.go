package ports

import "github.com/dimixlol/knowyourwebsite/domains/persister/ports"

type CacheManager interface {
	GetUrlBySlug(slug string) ports.URL
}

type Requester interface {
	MakeRequest() (string, error)
}
