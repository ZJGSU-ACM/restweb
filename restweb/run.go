package main

import (
	"log"
	"os"
	"os/exec"
)

func runApp() {
	if buildApp() != nil {
		return
	}

	os.Chdir(os.Getenv("GOPATH") + "/bin")

	cmd := exec.Command("./" + appName)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}
