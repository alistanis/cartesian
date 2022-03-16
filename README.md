# cartesian

Package cartesian implements a simple generic cartesian product in Go 1.18.

This is certainly not the fastest possible implementation.

https://en.wikipedia.org/wiki/Cartesian_product

To use:
```go
import "github.com/alistanis/cartesian"

func main() {
	i := []int{1, 2, 3} // or []float{3.14, ...} or []string{"hello", ...}
	j := []int{3, 1, 2}
	// Product accepts ...[]T any
	products := cartesian.Product(i, j) // [][]T
	
	for _, p := range products {
		fmt.Println(p)    
	}   
}

```

```
$ go run
[1 3]
[1 1]
[1 2]
[2 3]
[2 1]
[2 2]
[3 3]
[3 1]
[3 2]
```

Tests could be improved, and I will happily accept contributions.