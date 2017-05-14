# Panic After

Throws a panic after x seconds, useful for testing exception tracking and logging.

## Installation

Installation requires a working Go installation.

```
$ go get -u github.com/nicluo/panicafter
```

will download and install the `panicafter` command

## Usage

```
panicafter [seconds=10]
```

## Example

```
$ panicafter 1
panic: Oops

goroutine 1 [running]:
main.PanicAfter(0x1)
	.../go/src/github.com/nicluo/panicafter/panicafter.go:19 +0x79
main.main()
	.../go/src/github.com/nicluo/panicafter/panicafter.go:59 +0x7e
```
