# @Bithomp/xrpl-go-api A Bithomp Go library for XRPL

[![Go Report Card](https://goreportcard.com/badge/github.com/Bithomp/xrpl-go-api)](https://goreportcard.com/report/github.com/Bithomp/xrpl-go-api) [![GoDoc](https://pkg.go.dev/badge/github.com/Bithomp/xrpl-go-api?status.svg)](https://pkg.go.dev/github.com/Bithomp/xrpl-go-api)

# Tests

```Shell
go test -v ./...
```

# clean package

```Shell
go clean -i ./...

# To import all the necessary modules
go mod tidy
```

Generate address and seed

```Go
package main

import (
  "fmt"
  "github.com/Bithomp/xrpl-go-api/crypto"
  "github.com/Bithomp/xrpl-go-api/wallet"
)

func main() {
  // Generate an address and seed
  seed, address, _ := wallet.Generate(crypto.ALGORITHM_ED25519)
  fmt.Println("Address: ", address)
  fmt.Println("Seed: ", seed)
}
```

Convert classic address to X-address

```Go
package main

import (
  "fmt"
  "github.com/Bithomp/xrpl-go-api/address_codec"
)

func main() {
  // Convert classic address to X-address
  xAddress := address_codec.ClassicAddressToXAddress("rsuUjfWxrACCAwGQDsNeZUhpzXf1n1NK5Z", nil, false)
  fmt.Println("xAddress: ", xAddress) // X7czuu79XJ4GHhN5bsHDNyNjCrDFgjXw9rE9ELS86d47DXo
}
```
