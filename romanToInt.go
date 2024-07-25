package main

func romanToInt(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			num++
			break
		case 'V':
			if i != 0 {
				if s[i-1] == 'I' {
					num += 3
					break
				}

			}

			num += 5
			break
		case 'X':
			if i != 0 {
				if s[i-1] == 'I' {
					num += 8
					break
				}
			}
			num += 10
			break
		case 'L':
			if i != 0 {
				if s[i-1] == 'X' {
					num += 30
					break
				}
			}
			num += 50
			break
		case 'C':
			if i != 0 {
				if s[i-1] == 'X' {
					num += 80
					break
				}
			}
			num += 100
			break
		case 'D':
			if i != 0 {
				if s[i-1] == 'C' {
					num += 300
					break
				}
			}
			num += 500
			break
		case 'M':
			if i != 0 {
				if s[i-1] == 'C' {
					num += 800
					break
				}
			}
			num += 1000
			break
		}

	}
	return num
}
