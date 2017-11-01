package main

import (
	"fmt"

	"github.com/skrelan/LogrusWrapper/log"
)

func runTests() {
	fmt.Println("\nThis is a test \U00002699")
	log.Debug("Debug")
	log.Info("Info")
	log.Warn("Warn")
	log.Error("Error")
	log.AlwaysLog("AlwaysLog")
}

func main() {
	runTests()
	log.SetLogLevel("error")
	runTests()
}
