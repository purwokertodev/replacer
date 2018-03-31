## Replacer

Replace space character as defined by Unicode's White Space property (https://golang.org/pkg/unicode/#IsSpace)

[![Build Status](https://travis-ci.org/wuriyanto48/replacer.svg?branch=master)](https://travis-ci.org/wuriyanto48/replacer) [![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

### Usage

  - Install
    ```shell
    go get github.com/wuriyanto48/replacer
    ```
  - Example
    ```go
    package main

    import (
      "fmt"
      "github.com/wuriyanto48/replacer"
    )

    func main() {

      stringWithStrangeSpace := "sometimes i feel \n my heart \t so lonely"

      r1 := replacer.Replace(stringWithStrangeSpace, " ")
      r2 := replacer.Replace(stringWithStrangeSpace, "_")

      fmt.Println(r1) // sometimes i feel my heart so lonely
      fmt.Println(r2) // sometimes_i_feel_my_heart_so_lonely
    }
    ```
