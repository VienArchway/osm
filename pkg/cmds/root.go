package cmds

import (
	"flag"

	v "github.com/appscode/go/version"
	cfgCmd "github.com/appscode/osm/pkg/cmds/config"
	"github.com/spf13/cobra"
)

func NewCmdOsm() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "osm [command]",
		Short: `Object Store Manipulator by AppsCode`,
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	rootCmd.AddCommand(cfgCmd.NewCmdConfig())

	rootCmd.AddCommand(NewCmdListContainers())
	rootCmd.AddCommand(NewCmdMakeContainer())
	rootCmd.AddCommand(NewCmdRemoveContainer())

	rootCmd.AddCommand(NewCmdListIetms())
	rootCmd.AddCommand(NewCmdPush())
	rootCmd.AddCommand(NewCmdPull())
	rootCmd.AddCommand(NewCmdStat())
	rootCmd.AddCommand(NewCmdRemove())

	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}