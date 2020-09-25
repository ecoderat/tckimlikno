# tckimlikno

The tckimlikno package provides an validation for Turkish Republic's citizen ID numbers. The Id number is called "T.C. Kimlik No"


### Installation

```bash
$ go get -u github.com/ecoderat/tckimlikno
```


### Usage

```go
package main

import (
	"fmt"

	"github.com/ecoderat/tckimlikno"
)

func main() {
    b := tckimlikno.IsValid("01234567890")
    if b {
        fmt.Println("is a valid number")      
    } else {
        fmt.Println("is an invalid number")
    }
}
```
