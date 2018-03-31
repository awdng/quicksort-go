package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"time"
)

func main() {
	content, _ := ioutil.ReadFile("QuickSort.txt")

	var sort []int
	for _, v := range strings.Split(string(content), "\n") {
		i, err := strconv.Atoi(v)
		if err == nil {
			sort = append(sort, i)
		}
	}
	start := time.Now()
	quickSort(sort, 0, len(sort)-1)
	elapsed := time.Now().Sub(start)
	fmt.Print(sort)
	fmt.Print("\nTook: ", elapsed)
}

func partition(ar []int, p int, r int) int {
	x := ar[r]
	i := p - 1
	for j := p; j < r; j++ {
		if ar[j] <= x {
			i++
			ar[i], ar[j] = ar[j], ar[i]
		}
	}
	i++
	ar[i], ar[r] = ar[r], ar[i]
	return i
}

func quickSort(ar []int, p int, r int) {
	var q int
	if p < r {
		q = partition(ar, p, r)
		quickSort(ar, p, q - 1)
		quickSort(ar, q + 1, r)
	}
}