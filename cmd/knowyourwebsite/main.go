package main

import (
	"context"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister"
	"github.com/dimixlol/knowyourwebsite/domains/proxier"
	"github.com/dimixlol/knowyourwebsite/utils"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var (
	Version      = "0.0.1"
	rootCmd      = &cobra.Command{}
	persisterCmd = &cobra.Command{
		Use: "persister",
	}
	requesterCmd = &cobra.Command{
		Use: "proxier",
	}
)

func newListenerCmd(ctx context.Context, server *http.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "listen",
		Short: "listen",
		Long:  "listen long",
		Run: func(cmd *cobra.Command, args []string) {
			err := server.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}}
}

func main() {
	config.CreateConfiguration("config.json", Version)
	gin.SetMode(gin.ReleaseMode)
	tonic.SetErrorHook(utils.ErrHook)
	tonic.SetRenderHook(utils.RenderHook, "application/json")
	tonic.SetBindHook(utils.BindingHook)
	rootCtx := context.Background()

	persisterCmd.AddCommand(newListenerCmd(rootCtx, persister.NewHTTPPersister(rootCtx)))
	requesterCmd.AddCommand(newListenerCmd(rootCtx, proxier.NewHTTPRequester(rootCtx)))
	rootCmd.AddCommand(persisterCmd)
	rootCmd.AddCommand(requesterCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
