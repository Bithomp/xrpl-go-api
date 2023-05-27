# @Bithomp/xrpl-go-api

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
