package main

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister"
	"github.com/dimixlol/knowyourwebsite/domains/proxier"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/dimixlol/knowyourwebsite/utils"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "c", "config.json", "path to config file")
}

var (
	Version      = "0.0.1"
	rootCmd      = &cobra.Command{}
	persisterCmd = &cobra.Command{
		Use: "persister",
	}
	requesterCmd = &cobra.Command{
		Use: "proxier",
	}
	cfgFile string
)

func newListenerCmd(ctx context.Context, app string, server *http.Server) *cobra.Command {
	return &cobra.Command{
		Use: "listen",
		Run: func(cmd *cobra.Command, args []string) {
			logger := logging.GetLogger(ctx)
			logger.Infof(ctx, "Starting `%s` server version `%s`", app, Version)
			err := server.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}}
}

func main() {
	config.CreateConfiguration(cfgFile, Version)
	gin.SetMode(gin.ReleaseMode)
	tonic.SetErrorHook(utils.ErrHook)
	tonic.SetRenderHook(utils.RenderHook, "application/json")
	tonic.SetBindHook(utils.BindingHook)
	rootCtx := context.Background()

	persisterCmd.AddCommand(newListenerCmd(rootCtx, persisterCmd.Use, persister.NewHTTPPersister(rootCtx)))
	requesterCmd.AddCommand(newListenerCmd(rootCtx, requesterCmd.Use, proxier.NewHTTPRequester(rootCtx)))
	rootCmd.AddCommand(persisterCmd)
	rootCmd.AddCommand(requesterCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
