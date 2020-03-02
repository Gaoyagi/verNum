package main

import (
	"fmt"
	"os/exec"
)

//"sw_vers -productVersion" for current macOS   python3 --version for current python3version

func main() {
	//grab python3 version number
	pyVer, err1 := exec.Command("python3", "--version").Output()
	errorCatch(err1)

	//grab mac OS version number
	osVer, err2 := exec.Command("sw_vers",  "-productVersion").Output()
	errorCatch(err2)

	//grab mongoDB versions number
	mongodVer, err3 := exec.Command("mongod",  "--version").Output()
	errorCatch(err3)

	goVer, err4 := exec.Command("go",  "version").Output()
	errorCatch(err4)


	//prints out the desired version numbers
	fmt.Printf("Python3 current version: %s", string(pyVer)) 
	fmt.Printf("Mac OS current version: %s", string(osVer))
	fmt.Printf("MongoDB current version: %s", string(mongodVer))
	fmt.Printf("Golang current version: %s", string(goVer))
}

func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}
