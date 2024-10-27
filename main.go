package main

// Use: go get -u github.com/Lozlof/lozgo to update to the newest lozgo repo.
import (
	"fmt"
	"os"
	"log"
)

func main() {	
	// Display the welcome message.
	displayToTerminal(1)
	// Display the main menu.
	displayToTerminal(2)
}

// This function prints to the terminal.
// passedChoice is used to determine what to print.
func displayToTerminal(passedChoice int) {
	if passedChoice == 1 { // Will display the welcome message.
		// ReadFile reads the named file and returns the contents.
		// A successful call returns err == nil, not err == EOF.
		// Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported. 
		// EOF stands for End of File.
		// It’s an indicator in programming that signals when there is no more data to read from a file or data stream.
		// When you’re reading data (from a file, a socket, or another data source), the EOF condition tells you that you’ve reached the end.
		welcomeMessage, err := os.ReadFile("txtFiles/welcomeMessage.txt")
		
		if err != nil {
			log.Fatal(err)
		}
		
		// os.ReadFile stores the data as a byte slice in Go.
		// A byte slice in Go is essentially a slice of bytes, where each byte is an unsigned 8-bit integer (uint8).
		// Convert byte slice to string to display Unicode properly.
		// Without the string conversion it will not display correctly.
		fmt.Println(string(welcomeMessage))
	}

	if passedChoice == 2 { // Will display the main menu.
		mainMenuText, err := os.ReadFile("txtFiles/mainMenu.txt")

		if err != nil {
			log.Fatal(err)
		}

		// Will also not display correctly without the string conversion.
		fmt.Println(string(mainMenuText))
	}
}

// This function is used when user input is needed.
// inputChoice is used to determine what to execute.
func getUserInput(inputChoice int) {
	// userInput has to be a string because the lozgo sanitization.go functions require a string to be passed.
	// var userInput string

	if inputChoice == 1 {
		
		// can scans text read from standard input, storing successive space-separated values into successive arguments.
		// Newlines count as space.
		// It returns the number of items successfully scanned.
		// If that is less than the number of arguments, err will report why. 
		fmt.Scan(&userInput)
	}
}