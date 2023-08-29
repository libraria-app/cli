/*
Copyright © 2023 libraria-app
*/
package print

import "fmt"

const (
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
)

func Info(s string) {
	fmt.Println(colorGreen, s)
}

func Error(s string) {
	fmt.Println(colorRed, s)
}
