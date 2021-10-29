package main

import (
	"bufio"
	"fmt"
	"inccCalc/calc"
	"os"
	"strings"
)

func scan() {
	scanner := bufio.NewScanner(os.Stdin)
	keepReading := true
	for keepReading {
		fmt.Printf("inccCalc: ")
		if !(scanner.Scan()) {
			break
		}
		command := scanner.Text()
		keepReading = interpretCommand(command)
	}
}

func comHelp(command string) {
	switch command {
	case "calc":
		fmt.Println("--------------------------------------------------")
		fmt.Println("'calc [value] [startDate] (endDate)'")
		fmt.Println("All dates must be in the dd/mm/yyyy format.")
		fmt.Println("No separation characters are needed for the dates.")
		fmt.Println("endDate is optional: default is today.")
		fmt.Println("--------------------------------------------------")
	case "stop":
		fmt.Println("------------------")
		fmt.Println("'stop'")
		fmt.Println("Halts the program.")
		fmt.Println("------------------")
	case "help":
		fmt.Println("---------------------------------------------")
		fmt.Println("'help [command]'")
		fmt.Println("Gives a short description of a given command.")
		fmt.Println("-help")
		fmt.Println("-stop")
		fmt.Println("-version")
		fmt.Println("-calc")
		fmt.Println("---------------------------------------------")
	case "update":
		fmt.Println("WIP")
	default:
		fmt.Println("help " + command + ": '" + command + "' is not a known command. ")
	}
}

func printVersion() {
	fmt.Println("inccCalc version 1.0.")
	fmt.Println("Type 'help' for help, 'stop' to end the program.")
	fmt.Println("------------------------------------------------")
}

func interpretCommand(rawCommand string) bool {
	packedCommand := strings.Fields(rawCommand)
	switch packedCommand[0] {
	case "calc":
		calc.ComCalc(packedCommand)
	case "help":
		if len(packedCommand) == 1 {
			comHelp("help")
		} else {
			comHelp(packedCommand[1])
		}
	case "version":
		printVersion()
	case "stop":
		fmt.Println("Ending program...")
		return false
	default:
		fmt.Println("'" + packedCommand[0] + "' is not a known command. ")
	}
	return true

}

func main() {
	printVersion()
	scan()
}
