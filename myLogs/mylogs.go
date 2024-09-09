package myLogs

import (
	"log"
)

func MySuccessLog(message string) {
	green := "\033[32m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", green, message, reset)
}

func MyErrorLog(message string) {
	red := "\033[31m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", red, message, reset)
}

func MyInfoLog(message string) {
	blue := "\033[34m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", blue, message, reset)
}

func MyWarningLog(message string) {
	orange := "\033[33m"
	reset := "\033[0m"
	log.Printf("%s%s%s\n", orange, message, reset)
}
