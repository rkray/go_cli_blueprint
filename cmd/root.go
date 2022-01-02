// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:     "go_cli_blueprint",
    Short:   "GO CLI Blueprint",
    Long:    `Just a blueprint for small CLI application.
              Add some functionality.`,
    Run:     func(cmd *cobra.Command, args []string) {
        general()
    },
}

var Verbose bool
var Quiet bool

func Exec(version string) {
    rootCmd.Version="1.2.3"
    // global flags
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "quiet",   "q", false, "suppress any output")
    // commands
    rootCmd.AddCommand(cmdList)
    err:= rootCmd.Execute()
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
}
