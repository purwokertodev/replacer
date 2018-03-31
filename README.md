## Replacer

Replace space character as defined by Unicode's White Space property (https://golang.org/pkg/unicode/#IsSpace)

### Usage

  - Install
    ```shell
    go get github.com/wuriyanto48/replacer
    ```

  ```go
  package main

  import (
    "fmt"
    "github.com/wuriyanto48/replacer"
  )

  func main() {

    stringWithStrangeSpace := "sometimes i fell \n my heart \t so lonely"

    r1 := replacer.Replace(stringWithStrangeSpace, " ")
    r2 := replacer.Replace(stringWithStrangeSpace, "_")

    fmt.Println(r1) // sometimes i fell my heart so lonely
    fmt.Println(r2) // sometimes_i_fell_my_heart_so_lonely
  }
  ```
