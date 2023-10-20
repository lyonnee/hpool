# HPool (Happy Pool)

HPool is a generic object pool library for Go that provides a simple way to manage and reuse objects, reducing the overhead of object creation and destruction. The library is built on top of Go's `sync.Pool` and adds support for generics.

## Features

- Generic Support: HPool allows you to create object pools for different types, catering to various needs.
- High Performance: HPool performs well in concurrent environments.
- Easy to Use: Implementing object pooling with HPool requires only a few lines of code.

## Installation

Using Go Modules for project management, you can install HPool using the following command:

```shell
go get github.com/lyonee/hpool@latest
```

## Usage & Example
```go
import (
	"fmt"

	"github.com/lyonnee/hpool"
)

type People struct {
	Name string
	Age  int
}

func main() {
	pool := hpool.New[*People](func() *People {
		return &People{}
	})

	p := pool.Get()
    p.Name = "lyon.nee"
    p.Age = 27
    fmt.Println(p)

	pool.Put(p)
}
```

## Benchmarks

HPool benchmark results:
```
goos: windows
goarch: amd64
pkg: github.com/lyonnee/hpool
cpu: Intel(R) Core(TM) i5-7400 CPU @ 3.00GHz
BenchmarkStdPool-4   	82471392	        15.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkHPool-4     	75010312	        15.75 ns/op	       0 B/op	       0 allocs/op
PASS
	github.com/lyonnee/hpool	coverage: 100.0% of statements
ok  	github.com/lyonnee/hpool	2.869s
```

## License
This project is licensed under the MIT License. See the LICENSE file for more details.