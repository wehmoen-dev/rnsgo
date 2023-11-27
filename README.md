# RNS.GO - A simple Go implementation of RNS to resolve Names and Addresses

## Introduction

RNS.GO is a simple Go implementation of RNS to resolve Names and Addresses.

## Installation

```bash
go get github.com/wehmoen/rnsgo
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/wehmoen/rnsgo"
)

func main() {
	client := rnsgo.NewClient()

	// Resolve a name to an address
	address, err := client.GetAddress("jihoz.ron")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("RNS Name %s resolve to address: %s\n", "jihoz.ron", address)
	
	// Resolve an address to a name
	name, err := client.GetName("0xa09a9b6f90ab23fcdcd6c3d087c1dfb65dddfb05")
	
	if err != nil {
        fmt.Println(err)
        return
    }
	
	fmt.Printf("RNS Address %s resolve to name: %s\n", "0xa09a9b6f90ab23fcdcd6c3d087c1dfb65dddfb05", name)
	
	// Batch resolve names to addressess
	getAddressBatchInput := []rnsgo.Name{"jihoz.ron", "dwi.ron", "wehmoen.ron"}
	
	getAddressBatchResult, err := client.GetAddressBatch(getAddressBatchInput)
	
	if err != nil {
        fmt.Println(err)
        return
    }
	
	for name, address := range getAddressBatchResult {
		fmt.Printf("RNS Name %s resolve to address: %s\n", name, address)
	}
	
	// Batch resolve addresses to names
	getNameBatchInput := []rnsgo.Address{"0xa09a9b6f90ab23fcdcd6c3d087c1dfb65dddfb05", "0x445ba6f9f553872fa9cdc14f5c0639365b39c140", "0x3759468f9fd589665c8affbe52414ef77f863f72"}
	
	getNameBatchResult, err := client.GetNameBatch(getNameBatchInput)
	
	if err != nil {
        fmt.Println(err)
        return
    }
	
	for address, name := range getNameBatchResult {
		fmt.Printf("Address %s resolve to RNS name: %s\n", address, name)
	}
}

```

## License

MIT License

## Support

For support, please contact wehmoen on the Axie Infinity Discord server: https://discord.gg/axie