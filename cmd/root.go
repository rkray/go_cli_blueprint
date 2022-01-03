// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    //"fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:     "go_cli_blueprint",
    Short:   "GO CLI Blueprint",
    Long:    `Just a blueprint for small CLI application.
              Add some functionality.`,
    ValidArgs: []string{"v","q"},
    Run:     func(cmd *cobra.Command, args []string) {
        general()
    },
}

var Verbose bool
var Quiet bool

func Exec(version_id string) {
    // set version string
    rootCmd.Version=version_id
    cobra.CheckErr(rootCmd.Execute())
}

func init(){
    // global flags
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
    rootCmd.PersistentFlags().BoolVarP(&Quiet,   "quiet",   "q", false, "suppress any output")
}
