# go_mxchecker

This is a simple mailbox checker implemented in golang.

This will  find out the given email address is valid or not.


For Using first get this package by go get github.com/gokultp/go_mxchecker

Then try this example code

```go

package main

import(
  "fmt"
  "github.com/gokultp/go_mxchecker"
)

func main(){
    status, err :=  go_mxchecker.VerifyEmail("tp.gokul@gmail.com")
    
    if err!= nil {
        panic(err)
    }
    fmt.Println(status)
}
```
