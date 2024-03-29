package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/fs"
	"strings"
)

type proxierConfiguration struct {
	Host string
	Port string
}

type persisterConfiguration struct {
	Host          string
	Port          string
	SessionSecret string
	Session       struct {
		Secret string
		TTL    int
	}
	CSRF struct {
		Secret string
		Domain string
	}
}

type configuration struct {
	SlugLength  uint
	Environment string
	Version     string
	Viper       *viper.Viper
	DB          *dbConfiguration
	Cache       *cacheConfiguration
	API         *apiConfiguration
	Logging     *loggingConfiguration
	Proxier     *proxierConfiguration
	Persister   *persisterConfiguration
	CORS        *cors.Config
}

var Configuration *configuration

func New(configFile, version string) {
	boostrap(version)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(*fs.PathError); !ok {
			panic(err)
		}
	}
	err = viper.Unmarshal(&Configuration)
	if err != nil {
		panic(err)
	}
}

func boostrap(version string) {
	viper.SetDefault("slugLength", 10)
	viper.SetDefault("environment", gin.DebugMode)
	viper.SetDefault("logging.level", "info")
	// DB
	viper.SetDefault("db.port", "2345")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.user", "hpdb")
	viper.SetDefault("db.password", "hpdb")
	viper.SetDefault("db.name", "hpdb")
	viper.SetDefault("db.sslmode", "disable")
	viper.SetDefault("db.timezone", "UTC")
	viper.SetDefault("db.driver", "postgres")
	// CACHE
	viper.SetDefault("cache.host", "localhost")
	viper.SetDefault("cache.port", "6379")
	viper.SetDefault("cache.password", "")
	viper.SetDefault("cache.db", 0)
	viper.SetDefault("cache.size", 10000)
	viper.SetDefault("cache.ttl", "24h")
	// API
	viper.SetDefault("api.title", "Hosts-proxy API")
	viper.SetDefault("api.description", "Manage mapping between domains and IPs")
	viper.SetDefault("api.contact.email", "john.doe@mail.com")
	viper.SetDefault("api.license.name", "AGPLv3")
	viper.SetDefault("api.license.url", "https://github.com/dimixlol/hosts-proxy/raw/master/LICENSE")
	viper.SetDefault("api.logo.url", "/api/redoc/logo.png")
	viper.SetDefault("api.logo.color", "#fff")
	viper.SetDefault("api.logo.href", "https://foo.bar")
	// Persister
	viper.SetDefault("persister.host", "localhost")
	viper.SetDefault("persister.port", "8081")
	viper.SetDefault("persister.session.secret", "hello-world")
	viper.SetDefault("persister.session.ttl", 3600)
	viper.SetDefault("persister.csrf.secret", "hello-world")
	viper.SetDefault("persister.csrf.domain", "localhost")
	// Proxier
	viper.SetDefault("proxier.host", "localhost")
	viper.SetDefault("proxier.port", "8082")
	// CORS
	viper.SetDefault("cors.allowOrigins", []string{"http://localhost:5173"})
	viper.SetDefault("cors.allowMethods", []string{"GET", "POST", "OPTIONS"})
	viper.SetDefault("cors.allowHeaders", []string{"x-csrf-token", "content-type"})
	viper.SetDefault("cors.allowCredentials", true)
	viper.SetDefault("cors.maxAge", "12h")

	// constant
	viper.Set("version", version)
	viper.Set("api.version", "v1")
}
