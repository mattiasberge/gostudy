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
	swaps := 0 // total swaps
	swap := 0 // if swap within loop cycle or not
	
	for i = 0; i < length; i++ {
		l = append(l, rand.Intn(99))
	}
	
	for i = 0; i < len(l); i++ {
		swap = 0
		for j = 0; j < len(l); j++ {
			// if on last element of inner loop, break
			if j == (len(l) - 1) {
				break
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

		// no change for entire run, looks like we're done
		if swap == 0 { 
			break
		}
		
	}
	fmt.Printf("array_length=%d, iterations=%d, swaps=%d\n", len(l), iterations, swaps)
}