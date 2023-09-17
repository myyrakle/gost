package gost

import (
	"fmt"
	"unicode/utf8"
)

func Println(format String, params ...any) {
	Print(format, params...)
	fmt.Println()
}

func Print(format String, params ...any) {
	fmt.Print(Format(format, params...))
}

func Format(format String, params ...any) String {
	runes := []rune{}

	for len(string(format)) > 0 {
		r, size := utf8.DecodeRuneInString(string(format))
		runes = append(runes, r)
		format = format[size:]
	}

	paramsIndex := 0

	buffer := ""
	inLeftCurlyBrackets := false
	inRightCurlyBrackets := false

	for _, r := range runes {
		if r == '{' {
			if inLeftCurlyBrackets {
				buffer += "{"
				inLeftCurlyBrackets = false
				continue
			} else {
				inLeftCurlyBrackets = true
				continue
			}
		}

		if r == '}' {
			if inRightCurlyBrackets {
				buffer += "}"
				inRightCurlyBrackets = false
				continue
			} else if inLeftCurlyBrackets {
				param := params[paramsIndex]

				display := castToDisplay(param).Unwrap()
				buffer += string(display.Display())

				paramsIndex++
				inLeftCurlyBrackets = false
				continue
			} else {
				inRightCurlyBrackets = true
				continue
			}
		}

		buffer += string(r)
	}

	return String(buffer)
}
