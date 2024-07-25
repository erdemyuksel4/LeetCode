package main

func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}
	var pre string = string(strs[0][0:1])
	for i, a := 1, 0; len(strs[a]) >= i; a++ {
		if string(strs[a][0:i]) != pre {
			return pre[0 : len(pre)-1]
		}
		if a == len(strs)-1 {

			i++
			if len(strs[0]) < i {
				return pre
			}
			pre = string(strs[0][0:i])
			a = -1
		}

	}
	return pre[0 : len(pre)-1]
}
