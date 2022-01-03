// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
)

var setcnt int
var setname string

var cmdSet = & cobra.Command{
    Use:     "set",
    Short:   "Set something",
    Long:    "Set a value to",
    ValidArgs: []string{"name","n"},
    Deprecated: "do not use it anymore",
    Run:     func(cmd *cobra.Command, args []string) {
                fmt.Println(" => ValidateArgs:",cmd.ValidateArgs([]string{"name"}))
                set()
    },
}

func set(){
    fmt.Println("Set a value:", setname, "to", setcnt)

    if Quiet == true {
        fmt.Println("Quiet Mode On")
    }
    if Verbose == true {
        fmt.Println("Verbose Mode On")
    }
}

func init(){
    cmdSet.Flags().IntVarP(     &setcnt, "count", "c", 0, "example for int value")
    cmdSet.Flags().StringVarP( &setname,  "name", "n",  "noname", "example for string value")
    cmdSet.Flags().BoolP(  "force", "f",    false, "example for bool value")
    // command
    rootCmd.AddCommand(cmdSet)
}
