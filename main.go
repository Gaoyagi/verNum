package main

import (
	"fmt"
	"os/exec"
)

//"sw_vers -productVersion" for current macOS   python3 --version for current python3version

func main() {
	path, err := exec.LookPath("Python")
	errorCatch(err)
	//osVer, err := exec.Command("sw_vers -productVersion").Output()
	pyVer, err1 := exec.Command("/usr/bin/Python python3 --version", path).Output()
	errorCatch(err1)
	fmt.Printf("%s", pyVer)//"python3 version: %s\n", out)
	
}


func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}
