package mrctech

import (
	"sort"
)

func MiniProduct(a []int) int {
	sort.Ints(a)

	hasZero := false
	var posA []int
	var negA []int
	for _, v := range a {
		if v > 0 {
			posA = append(posA, v)
		} else if v < 0 {
			negA = append(negA, v)
		} else {
			hasZero = true
		}
	}

	sort.Ints(posA)
	sort.Ints(negA)
	//fmt.Println(posA)
	//fmt.Println(negA)
	negSize := len(negA)

	result := 1
	if negSize > 0 && negSize%2 == 0 {
		for i := 0; i < negSize-1; i++ {
			result = result * negA[i]
		}
		for _, v := range posA {
			result = result * v
		}
	} else if negSize > 0 {
		for _, v := range negA {
			result = result * v
		}
		for _, v := range posA {
			result = result * v
		}
	} else {
		if hasZero {
			result = 0
		} else {
			result = posA[0]
		}
	}

	return result
}
