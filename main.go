package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"os"
)

//"sw_vers -productVersion" for current macOS   python3 --version for current python3version

type Command []struct {
	Executable string   `json:"executable"`
	Commands []string 	`json:"commands"`
}


func main() {
	deleteFile()
	executes := openCommands()	//gets the json struct
	//loops through Command struct
	for x:=0; x<len(executes);  x++ {
		writeFile(executes[x].Executable +  ":")
		//loops through all the commands for executable
		for y:=0; y<len(executes[x].Commands); y++{
			temp := strings.Fields(executes[x].Commands[y])	//parses command string into words
			fmt.Println(temp)
			//checks of the command is a "where" command
			if temp[0] == "where" {
				path, err := exec.LookPath(temp[1])
				errorCatch(err)
				writeFile(executes[x].Commands[y] + ": " + path)	
			} else{		//other wise run the bash command as normal
				cmd, err := exec.Command(temp[0], temp[1]).Output()
				errorCatch(err)
				temp := string(cmd)
				temp  = strings.TrimSuffix(temp, "\n")
				writeFile(executes[x].Commands[y] + ": " + temp)
			}
		}
		writeFile("\n")
	}
}

//error thrower
func errorCatch(err error) {
	if err != nil {
		panic(err)
	}
}

//opens the json file and sotres the content in the Command struct
func openCommands() Command {
	//open file and error check it
	fileContents, err := ioutil.ReadFile("commands.json")
	errorCatch((err))
	//creates a struct to hold json information and reads the json file into it
	temp := Command{}
	json.Unmarshal(fileContents, &temp)
	return temp
}

//writes the version numbers to the text file
func writeFile(fileContents string) {
	file, err := os.OpenFile("versions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errorCatch((err))

	defer file.Close()

	if _, err = file.WriteString(fileContents + "\n"); err != nil {
		panic(err)
	}
}

//clears the text file 
func deleteFile() {
	if _, err := os.Stat("versions.txt"); os.IsNotExist(err) {	return }
	// delete file
	var err = os.Remove("versions.txt")
    errorCatch(err)
}

	
