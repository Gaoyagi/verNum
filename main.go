package main

import (
	"fmt"
	"os/exec"
)

//"sw_vers -productVersion" for current macOS   python3 --version for current python3version

func main() {
	pyVer, err1 := exec.Command("python3", "--version").Output()
	errorCatch(err1)
	
	osVer, err2 := exec.Command("sw_vers",  "-productVersion").Output()
	errorCatch(err2)

	fmt.Printf("current python3 version: %s", string(pyVer)) 
	fmt.Printf("current Mac OS version: %s", string(osVer))

}

func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}
