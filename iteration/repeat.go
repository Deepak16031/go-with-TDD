package iteration

import "strings"

func Repeat(character string, frequency int) string {
	var b strings.Builder
	for i := 0; i < frequency; i++ {
		b.WriteString(character)
	}
	return b.String()
}
