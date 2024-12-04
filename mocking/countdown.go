package mocking

import (
	"fmt"
	"io"
)

const countdownStart = 3

func Countdown(w io.Writer) {

	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(w, i)
	}
}
