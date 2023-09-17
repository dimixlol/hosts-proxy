package config

import "time"

type (
	dbConfiguration struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
		Timezone string
		Driver   string
	}

	cacheConfiguration struct {
		Host     string
		Port     string
		Password string
		DB       int
		Size     int
		TTL      time.Duration
	}
)
