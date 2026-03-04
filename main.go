package main

import (
	"fmt"

	zipfian "github.com/oryankibandi/zipf/zipf"
)

// MergeSortUint64 sorts a slice of uint64 using merge sort.
func MergeSortUint64(arr []uint64) {
	if len(arr) <= 1 {
		return
	}

	aux := make([]uint64, len(arr))

	var sort func(left, right int)
	sort = func(left, right int) {
		if left >= right {
			return
		}

		mid := left + (right-left)/2
		sort(left, mid)
		sort(mid+1, right)

		// Copy current window into auxiliary slice
		for i := left; i <= right; i++ {
			aux[i] = arr[i]
		}

		i := left
		j := mid + 1
		k := left

		for i <= mid && j <= right {
			if aux[i] >= aux[j] {
				arr[k] = aux[i]
				i++
			} else {
				arr[k] = aux[j]
				j++
			}
			k++
		}

		// Copy any remaining left-half elements
		for i <= mid {
			arr[k] = aux[i]
			i++
			k++
		}
	}

	sort(0, len(arr)-1)
}

func getZipfValues(exponent float64, offset float64, imax float64, itemCount uint64) {
	z := zipfian.NewZipf(exponent, offset, imax)
	if z == nil {
		panic("Unable to create Zipf generator")
	}

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("Zipf Exponent: %f\tImax: %f\n", exponent, imax)
	fmt.Println("-------------------------------------------------------------")
	items := make([]uint64, itemCount)
	for i := range itemCount {
		k := z.GetNext()
		// fmt.Printf("%d: %d\n", i, k)
		items[i] = k
	}
	// sort
	MergeSortUint64(items)
	for i, v := range items {
		if i == 0 {
			fmt.Printf("[ %d", v)
		} else if i == len(items)-1 {
			fmt.Printf(", %d]\n", v)
		} else {
			fmt.Printf(", %d", v)
		}
	}
	// fmt.Println("Items: ", items)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}

func main() {
	tests := []struct {
		s     float64
		v     float64
		imax  float64
		count uint64
	}{
		{s: 2, v: 1, imax: 100, count: 200},
		{s: 1.5, v: 1.2, imax: 80, count: 100},
		{s: 1.8, v: 5, imax: 250, count: 500},
		{s: 2.3, v: 20, imax: 1500, count: 1000},
	}

	for _, t := range tests {
		getZipfValues(t.s, t.v, t.imax, t.count)
	}
}
