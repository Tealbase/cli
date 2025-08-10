package cmd

import (
	"os"
	"os/signal"

	env "github.com/Netflix/go-env"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tealbase/cli/internal/status"
	"github.com/tealbase/cli/internal/utils"
)

var (
	override []string
	names    status.CustomName
	output   = utils.EnumFlag{
		Allowed: append([]string{utils.OutputEnv}, utils.OutputDefaultAllowed...),
		Value:   utils.OutputPretty,
	}
	statusCmd = &cobra.Command{
		GroupID: groupLocalDev,
		Use:     "status",
		Short:   "Show status of local Tealbase containers",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			es, err := env.EnvironToEnvSet(override)
			if err != nil {
				return err
			}
			return env.Unmarshal(es, &names)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, _ := signal.NotifyContext(cmd.Context(), os.Interrupt)
			return status.Run(ctx, names, output.Value, afero.NewOsFs())
		},
		Example: `  tealbase status -o env --override-name api.url=NEXT_PUBLIC_TEALBASE_URL
  tealbase status -o json`,
	}
)

func init() {
	flags := statusCmd.Flags()
	flags.VarP(&output, "output", "o", "Output format of status variables.")
	flags.StringSliceVar(&override, "override-name", []string{}, "Override specific variable names.")
	rootCmd.AddCommand(statusCmd)
}
