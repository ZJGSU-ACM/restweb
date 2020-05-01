package main

import (
	// "log"
	"os"
)

const AppConf = `
{
	"port":":8080",
	"sesson":true,
	"log":"Dev"
}
`

func newApp() {
	os.Mkdir(packageName, 0777)
	os.Mkdir(packageName+"/controller", 0777)
	os.Mkdir(packageName+"/views", 0777)
	os.Mkdir(packageName+"/config", 0777)
	appConf, err := os.Create(packageName + "/config/app.conf")
	if err != nil {
		appConf.Write([]byte(AppConf))
	}
	os.Mkdir(packageName+"/static", 0777)
}
