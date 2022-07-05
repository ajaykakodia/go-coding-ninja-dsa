package main

func replaceCharInString(str string, a, b string) string {
	if len(str) == 0 {
		return ""
	}
	smallString := replaceCharInString(str[1:], a, b)
	if string(str[0]) == a {
		return b + smallString
	}
	return string(str[0]) + smallString
}
