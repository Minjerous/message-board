package tool

import "fmt"

var sensitiveWords = make([]string, 0)

func CheckIfSensitive(s string) bool {
	sensitiveWords = append(sensitiveWords, "你妈", "傻逼", "sb", "垃圾", "神经病", "lj", "郭楠")
	for _, word := range sensitiveWords {
		if hashMatchFunc(s, word) {
			return true
		}
	}
	return false
}
func hash(str string, m []int) int {
	if len(str) == 0 {
		return 0
	}
	var (
		t   int
		res int = 0
	)
	for i := 0; i < len(str); i++ {
		t = m[i] * int(str[i]-'a')
		res = res + t
	}
	return res
}

func hashMatchFunc(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	}
	var m []int
	var t = 1
	m = append(m, 1)

	for i := 1; i < len(str2)+1; i++ {
		t = t * 26
		m = append(m, t) // m store with 26^0, 26^1, 26^2 ... 26^(len(str2))
	}

	str2Hash := hash(str2, m)
	fmt.Println(str2Hash)
	str1Hash := hash(string([]byte(str1)[:len(str2)]), m)

	if str2Hash == str1Hash {
		return true
	}

	for i := 1; i < len(str1)-len(str2)+1; i++ {
		newHash := (str1Hash-int(str1[i-1]-'a'))/26 +
			m[len(str2)-1]*int(str1[i+len(str2)-1]-'a')

		if newHash == str2Hash {
			return true
		} else {
			str1Hash = newHash
		}
	}
	return false
}
