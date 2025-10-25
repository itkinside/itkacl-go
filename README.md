# itkacl-go

## Overview

itkacl-go is a Go library for interfacing with the ITKACL access control system.

## Installation

Before using this package, make sure the ITKACL library is installed on your system.

Refer to the ITKACL documentation for installation instructions.

Then install this package:

```bash
go get github.com/itkinside/itkacl-go
```

## Usage

At the start of your application, initialize and free the ITKACL context:

```go
import (
	"github.com/itkinside/itkacl-go"
)

func main() {
	itkacl.ItkaclInit()
	defer itkacl.ItkaclFree()

	// the rest of your program
}
```

This creates a global, reusable context for ITKACL queries.

## Performing Access Checks

To perform an ITKACL access check:

```go
access, err := itkacl.ItkaclCheck(g.Realm, remoteUser)
if err != nil {
	log.Printf("Could not perform ITKACL check on realm %s: %v", g.Realm, err)
	return false
}

if access == nil {
	return false
}

return *access
```
