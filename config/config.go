package config

import (
	"github.com/spf13/viper"
	"io/fs"
	"strings"
	"time"
)

type proxierConfiguration struct {
	Host string
	Port string
	TLS  bool
}
type configuration struct {
	Host       string
	Port       string
	SlugLength uint
	DB         *dbConfiguration
	Cache      *cacheConfiguration
	Version    string
	API        *apiConfiguration
	Viper      *viper.Viper
	Logging    *loggingConfiguration
	Proxier    *proxierConfiguration
}

var Configuration *configuration

func CreateConfiguration(configFile, version string) {
	cfg := newConfiguration(version)
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()
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

func newConfiguration(version string) *viper.Viper {
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
	cfg.SetDefault("logging.level", "info")
	cfg.SetDefault("api.title", "Knowyourwebsite API")
	cfg.SetDefault("api.description", "Manage mapping between domains and IPs")
	cfg.SetDefault("api.contact.name", "dimixlol")
	cfg.SetDefault("api.contact.email", "dmitriy.t@dmxlol.io")
	cfg.SetDefault("api.contact.url", "https://dmxlol.io")
	cfg.SetDefault("api.license.name", "GPLv3")
	cfg.SetDefault("api.license.url", "https://github.com/dimixlol/knowyourwebsite/raw/master/LICENSE")
	cfg.SetDefault("api.logo.url", "https://redocly.github.io/redoc/petstore-logo.png")
	cfg.SetDefault("api.logo.color", "#fff")
	cfg.SetDefault("api.logo.altText", "Dimix Logo")
	cfg.SetDefault("api.logo.href", "https://dmxlol.io")
	cfg.SetDefault("proxier.host", "localhost")
	cfg.SetDefault("proxier.port", "8081")
	cfg.SetDefault("proxier.tls", false)
	// constant
	cfg.Set("version", version)
	cfg.Set("api.version", "v1")
	return cfg
}
