// Author: René Kray <rene@kray.info>
// Date:   2022-01-02

package cmd
import (
    "fmt"
)

func general(){
    fmt.Println("Here comes the list")
    if Verbose == true {
        fmt.Println("Verbose Mode On")
    }
}
