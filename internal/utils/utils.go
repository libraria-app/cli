/*
Copyright © 2023 libraria-app
*/
package utils

import "fmt"

const (
	colorGreen = "\033[32m"
	colorRed   = "\033[31m"
)

func PrintInfo(s string) {
	fmt.Println(colorGreen, s)
}

func PrintError(s string) {
	fmt.Println(colorRed, s)
}
