package main

import (
	"context"
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/domains/persister"
	"github.com/dimixlol/hosts-proxy/domains/proxier"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/dimixlol/hosts-proxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "c", "config.json", "path to config file")
	config.New(cfgFile, Version)
	gin.SetMode(config.Configuration.Environment)
	tonic.SetErrorHook(utils.ErrHook)
	tonic.SetRenderHook(utils.RenderHook, "application/json")
	tonic.SetBindHook(utils.BindingHook)
	persisterCmd.AddCommand(newListenerCmd(persisterCmd.Use, persister.NewHTTPPersister))
	requesterCmd.AddCommand(newListenerCmd(requesterCmd.Use, proxier.NewHTTPRequester))
	rootCmd.AddCommand(persisterCmd)
	rootCmd.AddCommand(requesterCmd)
}

var (
	Version      = "0.0.1"
	rootCmd      = &cobra.Command{}
	persisterCmd = &cobra.Command{Use: "persister"}
	requesterCmd = &cobra.Command{Use: "proxier"}
	cfgFile      string
)

func newListenerCmd(app string, makeSrv func(ctx context.Context) *http.Server) *cobra.Command {
	return &cobra.Command{
		Use: "listen",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			srv := makeSrv(ctx)
			logger := logging.GetLogger(ctx)
			logger.Infof(ctx, "Starting `%s` server version `%s` at http://%s", app, Version, srv.Addr)
			err := srv.ListenAndServe()
			if err != nil {
				panic(err)
			}
		},
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
