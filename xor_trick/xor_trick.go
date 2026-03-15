// XOR trick for finding a(/two) missing element(s)
// https://florian.github.io//xor-trick/
package main

import (
	"fmt"
	"math"
)


func buildSet(n int) map[uint]any {
	newSet := map[uint]any{}
	for i := 1; i <= n; i++{
		newSet[uint(i)] = struct{}{}
	}
	return newSet
}

func buildSetWithMissing(origin map[uint]any, n, m int) map[uint]any {
	newSet := map[uint]any{}
	for k, v := range(origin) {
		newSet[k] = v
	}
	delete(newSet, uint(n))
	delete(newSet, uint(m))
	return newSet
}

func xorList(nums map[uint]any) uint {
	var sum uint
	for num := range nums {
		sum ^= num
	}
	return sum
}

func partition(nums map[uint]any, lsb uint) map[uint]any {
	newNums := map[uint]any{}
	for num := range nums {
		if (num >> lsb) & 1 == 1 {
			newNums[num] = struct{}{}
		}
	}
	return newNums
}

func main() {
	allNumbers := buildSet(36)
	numbersGiven := buildSetWithMissing(allNumbers, 27, 26)
	var results uint
	results = xorList(allNumbers) ^ xorList(numbersGiven)
	leastSignificantBit := uint(0)
	i := uint(0)
	for {
		if (results>>i) & 1 == 1 {
			leastSignificantBit = i
			break
		}
		i++
		if i >= uint(math.Pow(2, 32)) {
			break
		}
	}
	partitionAofAllNumbers := partition(allNumbers, leastSignificantBit)
	partitionAofNumbersGiven := partition(numbersGiven, leastSignificantBit)
	x0 := xorList(partitionAofAllNumbers) ^ xorList(partitionAofNumbersGiven)
	x1 := x0 ^ results
	fmt.Printf("Final results of numbers that are missing is: %d and %d\n", x0, x1)
}
