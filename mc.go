package main
import (
	"flag"    // used to parse cli-arguments
	"log"     // used to maintain logs
	"os"      // platform independent interface to os functions
	"os/exec" // runs external commands
)

func executeCommand(command string, argsArray []string) error { // function executeCommand uses a command and an array of arguments
	args := argsArray
	cmdObj := exec.Command(command, args...) // initialize and declare object
	cmdObj.Stdout = os.Stdout                // stdout to display the output to screen
	cmdObj.Stderr = os.Stderr                // stderr to handle errors
	cmdObj.Stdin = os.Stdin                  // stdin to add and input commands
	err := cmdObj.Run()						// Run the object

	if err != nil {
		log.Fatal(err)						//errors are logged
	}
	return nil
}

func main() {
	iface := flag.String("i", "eth0", "Interface name")				//flag for interface
	newmac := flag.String("m", "", "New MAC address")				//flag for new mac
	flag.Parse()													//parse the flags

	if *newmac == "" {
		log.Fatal("New MAC address must be specified using -m flag")	//error handling
	}

	command := "sudo"												
	executeCommand(command, []string{"ifconfig", *iface, "down"})	//interface down
	executeCommand(command, []string{"ifconfig", *iface, "hw", "ether", *newmac})	// interace's ether value changes to new mac
	executeCommand(command, []string{"ifconfig", *iface, "up"})		//interface up
}
//Mac addressed changes successfully