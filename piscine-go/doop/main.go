package main

import (
	"os"
)

func writeString(s string) {
	os.Stdout.WriteString(s)
}

func parseInt(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	neg := false
	i := 0
	if s[0] == '-' {
		neg = true
		i++
	} else if s[0] == '+' {
		i++
	}
	if i == len(s) {
		return 0, false
	}
	var n int64
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		digit := int64(c - '0')
		if neg {
			if n < (-9223372036854775808+digit)/10 {
				return 0, false
			}
			n = n*10 - digit
		} else {
			if n > (9223372036854775807-digit)/10 {
				return 0, false
			}
			n = n*10 + digit
		}
	}
	return n, true
}

func intToString(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
	}
	buf := [20]byte{}
	i := 20
	for n != 0 {
		i--
		digit := n % 10
		if digit < 0 {
			digit = -digit
		}
		buf[i] = byte(digit) + '0'
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, okA := parseInt(os.Args[1])
	op := os.Args[2]
	b, okB := parseInt(os.Args[3])

	if !okA || !okB {
		return
	}

	var result int64

	switch op {
	case "+":
		result = a + b
		if (b > 0 && result < a) || (b < 0 && result > a) {
			return
		}
	case "-":
		result = a - b
		if (b < 0 && result < a) || (b > 0 && result > a) {
			return
		}
	case "*":
		if a != 0 && b != 0 {
			result = a * b
			if result/a != b {
				return
			}
		}
	case "/":
		if b == 0 {
			writeString("No division by 0\n")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			writeString("No modulo by 0\n")
			return
		}
		result = a % b
	default:
		return
	}

	writeString(intToString(result) + "\n")
}
