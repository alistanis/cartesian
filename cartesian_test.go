package cartesian

import (
	"testing"
)

func TestProduct(t *testing.T) {
	{
		i := []int{1, 2, 3}
		j := []int{3, 2, 1}
		k := []int{2, 1, 3}

		prod := Product(i, j, k)
		n := len(prod)
		e := len(i) * len(j) * len(k)
		if n != e {
			t.Fatalf("there should be %d but there were %d groups\n", e, n)
		}

	}

	{
		i := []float64{1, 3.14, 42}
		j := []float64{42, 3.14, 1}
		k := []float64{3.14, 1, 42}

		prod := Product(i, j, k)
		n := len(prod)
		e := len(i) * len(j) * len(k)
		if n != e {
			t.Fatalf("there should be %d but there were %d groups\n", e, n)
		}
	}

	{
		i := []string{"hello", "goodbye", "42"}
		j := []string{"hello", "goodbye", "42"}
		k := []string{"hello", "goodbye", "42"}

		prod := Product(i, j, k)
		n := len(prod)
		e := len(i) * len(j) * len(k)
		if n != e {
			t.Fatalf("there should be %d but there were %d groups\n", e, n)
		}
	}

}
