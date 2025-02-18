package cli

import (
	gw_model "github.com/centralmind/gateway/model"
	"github.com/centralmind/gateway/restgenerator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func REST(configPath *string, addr *string) *cobra.Command {
	return &cobra.Command{
		Use:   "rest",
		Short: "REST gateway",
		Args:  cobra.MatchAll(cobra.ExactArgs(0)),
		RunE: func(cmd *cobra.Command, args []string) error {
			gwRaw, err := os.ReadFile(*configPath)
			if err != nil {
				return errors.Errorf("unable to read yaml config file: %w", err)
			}
			gw, err := gw_model.FromYaml(gwRaw)
			if err != nil {
				return errors.Errorf("unable to parse config file: %w", err)
			}
			mux := http.NewServeMux()
			a, err := restgenerator.NewAPI(*gw)
			if err != nil {
				return errors.Errorf("unable to init api: %w", err)
			}
			a.RegisterRoutes(mux)
			return http.ListenAndServe(*addr, mux)
		},
	}
}
