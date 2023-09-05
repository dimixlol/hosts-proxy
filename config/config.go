package config

import (
	"github.com/spf13/viper"
	"io/fs"
	"time"
)

type dbConfiguration struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	Timezone string
	Driver   string
}

type cacheConfiguration struct {
	Host     string
	Port     string
	Password string
	DB       int
	Size     int
	TTL      time.Duration
}

type configuration struct {
	Host       string
	Port       string
	SlugLength uint
	DB         *dbConfiguration
	Cache      *cacheConfiguration
	Viper      *viper.Viper
}

var Configuration *configuration

func CreateConfiguration(configFile string) {
	cfg := newConfiguration()
	cfg.SetConfigFile(configFile)
	err := cfg.ReadInConfig()
	if err != nil {
		if _, ok := err.(*fs.PathError); !ok {
			panic(err)
		}
	}
	err = cfg.Unmarshal(&Configuration)
	if err != nil {
		panic(err)
	}
}

func newConfiguration() *viper.Viper {
	cfg := viper.New()
	cfg.SetDefault("host", "")
	cfg.SetDefault("port", "8080")
	cfg.SetDefault("slugLength", 10)
	cfg.SetDefault("db.port", "2345")
	cfg.SetDefault("db.host", "localhost")
	cfg.SetDefault("db.user", "pq_local_app")
	cfg.SetDefault("db.password", "pq_local_app")
	cfg.SetDefault("db.name", "knowyourwebsite")
	cfg.SetDefault("db.sslmode", "disable")
	cfg.SetDefault("db.timezone", "UTC")
	cfg.SetDefault("db.driver", "postgres")
	cfg.SetDefault("cache.host", "localhost")
	cfg.SetDefault("cache.port", "6379")
	cfg.SetDefault("cache.password", "")
	cfg.SetDefault("cache.db", 0)
	cfg.SetDefault("cache.size", 10000)
	cfg.SetDefault("cache.ttl", 24*time.Hour)

	return cfg
}
