package main



func main() {
	s:="abcdef"
	reverseString([]byte(s))

}

func reverseString(s []byte) {

	l, r := 0, len(s)-1
	for l < r {
		var temp = s[l]
		s[l] = s[r]
		s[r] = temp
		l++
		r--
	}
}
