package iteration

import "strings"

func Repeat(stringToRepeat string, times int) (result string) {
	result = strings.Repeat(stringToRepeat, times)
	return
}

func RepeatWithFor(stringToRepeat string, times int) (result string) {
	for range times {
		result += stringToRepeat
	}
	return
}

func RepeatWithForUsingJoin(stringToRepeat string, times int) (result string) {
	for range times {
		strings.Join([]string{result, stringToRepeat}, "")
	}
	return
}
