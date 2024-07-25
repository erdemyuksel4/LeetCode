package main

func isPalindrome(x int) bool {
	var revInt int = 0
	cpyX := x
	for x > 0 {
		revInt = revInt*10 + x%10
		x /= 10
	}
	if cpyX == revInt {
		return true
	}
	return false
}
