package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	cmd := exec.Command("cd ..")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)

	out, err := exec.Command("cd ..").Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", out)//"python3 version: %s\n", out)
}
