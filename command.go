package main

import (
	"fmt"
	"github.com/Chinwendu20/otel_components_generator/config"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

// Command is the main entrypoint for this application
func command(cfg config.ConfigStruct) (*cobra.Command, error) {
	flagSet := flags()
	cmd := &cobra.Command{
		SilenceUsage:  true, // Don't print usage on Run error.
		SilenceErrors: true, // Don't print errors; main does it.
		Use:           "ocg",
		Long: fmt.Sprintf("OpenTelemetry Collector Builder (%s)", version) + `

ocg generates a custom OpenTelemetry Collector binary using the
options supplied by the commandline options. If no options are supplied
ocg requests for these interactively.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			checkEmptyConfigOptions(cfg)
			if err := validateComponent(Config); err != nil {

				return err
			}
			return generateComponent(cfg)
		},
	}

	cmd.Flags().AddGoFlagSet(flagSet)
	return cmd, nil
}
