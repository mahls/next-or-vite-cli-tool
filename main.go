package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Choose an option:")
	fmt.Println("1. Install Next.js app")
	fmt.Println("2. Install Vite app")

	var choice int
	fmt.Print("Enter your choice (1 or 2): ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("Enter the project name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	projectName := scanner.Text()

	switch choice {
	case 1:
		fmt.Printf("Installing Next.js app with project name '%s'...\n", projectName)
		err := runCommand("npx", "create-next-app", projectName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing Next.js: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Next.js installation complete.")
	case 2:
		fmt.Printf("Installing Vite app with project name '%s'...\n", projectName)
		err := runCommand("npm", "create", "vite", projectName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing Vite: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Vite installation complete.")
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

