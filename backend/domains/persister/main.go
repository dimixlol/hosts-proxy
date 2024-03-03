package persister

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimixlol/hosts-proxy/adapters/storage"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/domains/persister/entrypoints/api"
	"github.com/dimixlol/hosts-proxy/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
)

const (
	sessionName         = "__persister_sess_id__"
	csrfCookieName      = "__persister_csrf__"
	healthCheckEndpoint = "/ping"
)

var healthCheckResponse = map[string]string{"message": "pong"}

func NewHTTPPersister(ctx context.Context) *http.Server {
	engine := utils.NewApiEngine(constructOpenApiInfo, newSessionMiddleware(), newCsrfMiddleware(), cookieSetter)
	engine.GET(healthCheckEndpoint, nil, healthCheckHandler)
	db := storage.NewDatabaseStorage(ctx)
	api.NewHTTPAPI(ctx, db, engine)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Persister.Host, config.Configuration.Persister.Port),
		Handler: engine,
	}
}

func cookieSetter(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     csrfCookieName,
			Value:    csrf.GetToken(c),
			MaxAge:   config.Configuration.Persister.Session.TTL,
			Domain:   config.Configuration.Persister.CSRF.Domain,
			Path:     "/",
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
		})
	}
}

func newCsrfMiddleware() gin.HandlerFunc {
	return csrf.Middleware(csrf.Options{
		Secret: config.Configuration.Persister.CSRF.Secret,
		ErrorFunc: func(c *gin.Context) {
			c.AbortWithStatusJSON(
				utils.NewUnsuccessfulResponseWithCode(http.StatusForbidden, "invalid csrf token"),
			)
		},
		TokenGetter: func(c *gin.Context) string {
			token, err := c.Cookie(csrfCookieName)
			if errors.Is(err, http.ErrNoCookie) && gin.Mode() != gin.ReleaseMode {
				return c.Request.Header.Get("X-CSRF-TOKEN")
			}
			return token
		},
	})
}

func newSessionMiddleware() gin.HandlerFunc {
	return sessions.Sessions(
		sessionName,
		cookie.NewStore(
			[]byte(config.Configuration.Persister.Session.Secret),
		),
	)
}

func healthCheckHandler(c *gin.Context) {
	c.AbortWithStatusJSON(utils.NewSuccessfulResponseWithCode(http.StatusOK, healthCheckResponse))
}
