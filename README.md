# GoNetProbe

GoNetProbe is a simple, concurrent network scanner library written in Go. It provides functions to scan TCP ports on a given host.

## Installation

To install GoNetProbe, you can use `go get`:

```bash
go get github.com/t00mas/GoNetProbe
```

## Usage

Here's a simple example of how to use GoNetProbe:

```go
go package main

import (
    "fmt"
    "net"
    "time"
    "github.com/t00mas/GoNetProbe"
)

type dialer struct{}

func (d *dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
    return net.DialTimeout(network, address, timeout)
}

func main() {
    d := &dialer{}
    fmt.Println("Starting initial scan...")
    gonetprobe.InitialScan(d, "localhost")
    fmt.Println("Starting wide scan...")
    gonetprobe.WideScan(d, "localhost")
}
```

In this example, we define a simple `dialer` type that implements the `Dialer` interface required by GoNetProbe. We then pass an instance of this dialer to the `InitialScan` and `WideScan` functions.

## Testing

To run the tests for GoNetProbe, you can use `go test`:

```bash
go test github.com/t00mas/GoNetProbe
```

## Contributing

Contributions to GoNetProbe are welcome! Please submit a pull request or create an issue on GitHub.

## License

GoNetProbe is licensed under the MIT License. See the `LICENSE` file for details.
