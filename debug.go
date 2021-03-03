package main

import "fmt"

const debugMode = false

func debugPrintf(format string, args ...interface{}) {
	if debugMode {
		fmt.Printf(format, args...)
	}
}

func debugPrint(args ...interface{}) {
	if debugMode {
		fmt.Print(args...)
	}
}
func debugPrintln(args ...interface{}) {
	if debugMode {
		fmt.Println(args...)
	}
}
