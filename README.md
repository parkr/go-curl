# go-curl

A stripped-down, static alternative to `curl`. Useful for HEALTHCHECKs in Docker.

## Installation

```console
$ go get github.com/parkr/go-curl/...
$ go-curl -h
```

## Usage

```console
$ go-curl <url>
```

Want it to fail on non-200's? Pass the `-f` flag:

```console
$ go-curl -f <url>
```

## License

MIT. See `LICENSE` in this repository.
