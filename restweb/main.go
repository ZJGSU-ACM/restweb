package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: restweb [cmd] [app]")
}

var packageName string
var appName string

func getAppName() string {
	p := strings.Split(packageName, "/")
	return p[len(p)-1]
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	cmd := os.Args[1]
	packageName = os.Args[2]
	appName = getAppName()
	switch cmd {
	case "new":
		newApp()
	case "build":
		buildApp()
	case "run":
		runApp()
	case "clean":
		cleanApp()
	default:
		usage()
	}
}

func cleanApp() {
	os.Remove(packageName + "/" + packageName)
	os.Remove(packageName + "/main.go")
	os.Remove(packageName + "/config/router.conf")
}
