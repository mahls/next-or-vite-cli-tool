package main

import (
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

	switch choice {
	case 1:
		fmt.Println("Installing Next.js app...")
		err := runCommand("npx", "create-next-app", "my-next-app")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error installing Next.js: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Next.js installation complete.")
	case 2:
		fmt.Println("Installing Vite app...")
		err := runCommand("npm", "create", "vite", "my-vite-app")
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

