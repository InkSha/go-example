# Go Example

This is a simple Go example package.

## Installation

To install the package, run:

```bash
go get github.com/username/go-example
```

### Installation Commandline Tool

```bash
go install github.com/username/go-example/cmd/go-example

go-example
# output:
# current working directory:  /path/to/your/current/directory
```

## Usage

```go
package main

import (
  print "github.com/username/go-example/print"
)

func main() {
  print.ExamplePrint("hello world", 3)
}
```
