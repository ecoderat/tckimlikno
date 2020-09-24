package tckimlikno

import (
	"strconv"
)

// IsValid returns whether it is a valid number
func IsValid(tc string) bool {
	if tc[0] == '0' {
		return false
	}
	if len(tc) != 11 {
		return false
	}

	if isDigit(tc) != true {
		return false
	}

	arr := make([]int, 11)

	for i, j := range tc {
		arr[i], _ = strconv.Atoi(string(j))
	}

	var oddSum = arr[0]
	
	var evenSum int
	
	for i := 1; i < 8; i++ {
		evenSum += arr[i] 	//  arr[1] 3 5 7
		i++
		oddSum += arr[i]	// 	arr[2] 4 6 8

	}

	n := ((oddSum * 7) - evenSum) % 10
	if n < 0 {
		n += 10
	}
	if n != arr[9] {
		return false
	}

	m := (oddSum + evenSum + arr[9]) % 10 

	if m != arr[10] {
		return false
	}

	return true
}

// isDigit returns whether the given string consists of numbers only 
func isDigit(s string) bool {
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	return b
}
