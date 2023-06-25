package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	var arr []int
	if x < 0 {
		return false
	}
	for {
		if x <= 0 {
			break
		}
		arr = append(arr, x%10)
		fmt.Println(arr)
		x = x / 10
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

func reverseNumber(x int) int {
	temp := 0
	for {
		if x <= 0 {
			break
		}
		temp = temp*10 + x%10
		x = x / 10
	}
	return temp
}

func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	len := len(s)
	for i := 0; i < len/2; i++ {
		if s[i] != s[len-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-1221))
	fmt.Println(reverseNumber(-13324))
	x := 133221
	reverse := reverseNumber(x)
	if x == reverse {
		fmt.Println("Its a Palindrome")
	} else {
		fmt.Println("Its not a Palindrome")
	}
	fmt.Println("Check palindrome by string comparison ", isPalindromeString(123421))

}
