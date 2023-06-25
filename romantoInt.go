package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	fmt.Println(romanToIntImproved("MCMXCIV"))
}

func romanToInt(s string) int {
	var arr []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I': //converted to rune
			arr = append(arr, 1)
		case 'V':
			arr = append(arr, 5)
		case 'X':
			arr = append(arr, 10)
		case 'L':
			arr = append(arr, 50)
		case 'C':
			arr = append(arr, 100)
		case 'D':
			arr = append(arr, 500)
		case 'M':
			arr = append(arr, 1000)
		}
	}
	num := 0
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			num = num + arr[i]
			break
		}
		if arr[i] < arr[i+1] {
			num = num + (arr[i+1] - arr[i])
			i++
		} else {
			num = num + arr[i]
		}
	}
	return num
}

func romanToIntImproved(s string) int {
	//do it on the go

	//normal add or subtract two and add

	var ans int
	counter := 0

	for i, v := range s {
		if counter > i {
			continue
		}
		var temp1, temp2 int
		// fmt.Println(i, v) // v is the ASCII value of the string literal
		switch v {
		case 'I':
			temp1 = 1
		case 'V':
			temp1 = 5
		case 'X':
			temp1 = 10
		case 'L':
			temp1 = 50
		case 'C':
			temp1 = 100
		case 'D':
			temp1 = 500
		case rune('M'):
			temp1 = 1000
		}

		//check if it is last element
		if i != len(s)-1 {
			switch rune(s[i+1]) {
			case 'I':
				temp2 = 1
			case 'V':
				temp2 = 5
			case 'X':
				temp2 = 10
			case 'L':
				temp2 = 50
			case 'C':
				temp2 = 100
			case 'D':
				temp2 = 500
			case rune('M'):
				temp2 = 1000
			}
			if temp1 < temp2 {
				temp1 = temp2 - temp1
				counter = counter + 1
			}
		}
		ans = ans + temp1
		counter = counter + 1
	}
	return ans

}
