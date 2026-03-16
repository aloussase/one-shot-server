# one-shot-server

A simple one-shot server to use while testing integrations with 3rd party APIs.

## Usage

```bash
$ one-shot-server -help 
Usage: one-shot-server [OPTIONS]

  -body string
        A path to a file containing the resource to serve as JSON. May be omitted.
  -help
        Print usage information
  -path string
        The path from which to serve the resource (default "/")
  -port int
        The port from which to start the server (default 3000)
  -status int
        The status code with which to respond (default 200)
```

## Installation

### Go

```bash
go install github.com/aloussase/one-shot-server
```

### Docker

## License

MIT
