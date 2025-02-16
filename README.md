# TCPWrapper

TCPWrapper is a Go package that provides a middleware-enabled wrapper for TCP connections, allowing request and response processing through customizable middleware chains.

## Features

- Middleware support for request and response processing
- Automatic detection of HTTP responses
- Support for both Content-Length-based and delimiter-based message reading
- Graceful connection handling

## Installation

```sh
go get github.com/mmskazak/tcpwrapper
```

## Usage

### Creating a TCP Wrapper

```go
package main

import (
 "fmt"
 "net"
 "github.com/yourusername/tcpwrapper"
 "golang.org/x/text/encoding/charmap"
 "golang.org/x/text/transform"
 "bytes"
 "io"
)

// RequestLoggerMiddleware logs incoming requests and converts from Win1251 to UTF-8
func RequestLoggerMiddleware(data []byte) ([]byte, error) {
 fmt.Println("Received request:", string(data))

 reader := transform.NewReader(bytes.NewReader(data), charmap.Windows1251.NewDecoder())
 utf8Data, err := io.ReadAll(reader)
 if err != nil {
  return nil, err
 }

 return utf8Data, nil
}

// ResponseLoggerMiddleware logs response length
func ResponseLoggerMiddleware(data []byte) ([]byte, error) {
 fmt.Printf("Response length: %d bytes\n", len(data))
 return data, nil
}

func main() {
 listener, err := net.Listen("tcp", ":8080")
 if err != nil {
  panic(err)
 }
 defer listener.Close()

 for {
  conn, err := listener.Accept()
  if err != nil {
   fmt.Println("Error accepting connection:", err)
   continue
  }

  wrapper := &tcpwrapper.TCPWrapper{
   Conn:             conn,
   RequestDelimiter: []byte("\n"),
  }

  wrapper.AddRequestMiddleware(RequestLoggerMiddleware)
  wrapper.AddResponseMiddleware(ResponseLoggerMiddleware)

  go func() {
   defer wrapper.Close()
   if err := wrapper.HandleMessage(); err != nil {
    fmt.Println("Error handling message:", err)
   }
  }()
 }
}
```

## Using TCPWrapper with Caddy

To integrate TCPWrapper with Caddy, you need a custom module that wraps TCP connections. Below is an example configuration:

### Caddy Module Example

```json
{
  "apps": {
    "tcp": {
      "listeners": [
        {
          "address": ":8080",
          "handler": {
            "wrapper": "tcpwrapper",
            "request_delimiter": "\n"
          }
        }
      ]
    }
  }
}
```

### Running Caddy with TCPWrapper

Ensure you have the TCP module for Caddy installed and configure it to use the wrapper. You may need to build a custom Caddy version with your module.

```sh
caddy run --config caddy.json
```

Now, your Caddy server will process incoming TCP connections using TCPWrapper, applying middleware before forwarding the requests.

## License

This project is licensed under the MIT License.
