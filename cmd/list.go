// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
)

var cnt int
var name string
var force bool

var cmdList = & cobra.Command{
    Use:     "list",
    Aliases: []string{"ls","ll"},
    Short:   "List something",
    Long:    "Show a list with things",
    ValidArgs: []string{"long","l"},
    Run:     func(cmd *cobra.Command, args []string) {
        fmt.Println(" => ValidateArgs:",cmd.ValidateArgs([]string{"name"}))
        list()
    },
}

func list(){
    fmt.Println("Listname:", name)
    if cnt != 0 {
        for i:=0; i<=cnt; i++ {
            fmt.Println("Here comes the list", i)
        }
    } else {
        fmt.Println("Here comes the list")
    }
    if Quiet == true {
        fmt.Println("Quiet Mode On")
    }
    if Verbose == true {
        fmt.Println("Verbose Mode On")
    }
}

func init(){
    cmdList.Flags().IntVarP(     &cnt, "count", "c", 0, "example for int value")
    cmdList.Flags().StringVarP( &name,  "name", "n",  "noname", "example for string value")
    cmdList.Flags().BoolP(  "force", "f",    false, "example for bool value")
    // command
    rootCmd.AddCommand(cmdList)
}
