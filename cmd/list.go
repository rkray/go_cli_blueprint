// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
)

var cmdList = & cobra.Command{
    Use:     "list",
    Aliases: []string{"ls","ll"},
    Short:   "List something",
    Long:    "Show a list with things",
    //Args:    func(cmd *cobra.Command, args []string) error {
    //    
    //},
    Run:     func(cmd *cobra.Command, args []string) {
        list()
    },
}

func list(){
    fmt.Println("Here comes the list")
    if Verbose == true {
        fmt.Println("Verbose Mode On")
    }
}
