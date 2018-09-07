# Instagram-Go

Instagram-Go is a [Instagram API](https://www.instagram.com/developer/) client written in Go. Currently it can:

- Get Access Token (OAuth)
- Get Self

## Installation

```bash
$ go get github.com/swensonhe/instagram-go
```

## Examples

### Get Self

```go
package main

import (
    "github.com/swensonhe/instagram-go"
    "fmt"
)

func main() {
    client := instagram.NewClient(&instagram.Config{
    	ClientId: "CLIENT_ID",
        ClientSecret: "CLIENT_SECRET",
        RedirectUri: "https://www.swensonhe.com/",
    })
    res, err := client.GetSelf("token")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(res.Data)
}
```
