package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//TODO Get and check flags
	checkFlags()

	//TODO Create auth header

	//TODO Implement different methods
	//TODO Parse to valid XML using EBICS standard

	fmt.Println("EBICS client")
}

func checkFlags() {
	functionFlag := flag.String("run", "", "Function to run")

	flag.Parse()

	// Dereference
	funcParsed := *functionFlag

	switch funcParsed {
	case "generateKeys":

		// Generate client keys for authentication
		keyGen := generateKeyPair()
		if !keyGen {
			fmt.Println("Could not generate keys")
			os.Exit(1)
		}

		os.Exit(0)
		break
	}
}
