package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	length := 200
	var l []int
	var tmp int
	i := 0
	j := 0
	iterations := 0
	swaps := 0
	swap := 0
	
	for i = 0; i < length; i++ {
		l = append(l, rand.Intn(99))
	}
	
	for i = 0; i < len(l); i++ {
		swap = 0
		for j = 0; j < len(l); j++ {
			if j == (len(l) - 1) {
				continue
			}
			if l[j] > l[j+1] {
				tmp = l[j+1]
				l[j+1] = l[j]
				l[j] = tmp
				
				swaps++
				swap = 1
				
			}
			iterations++
		}
		if swap == 0 {
			break
		}
		
	}
	fmt.Printf("array_length=%d, iterations=%d, swaps=%d\n", len(l), iterations, swaps)
}