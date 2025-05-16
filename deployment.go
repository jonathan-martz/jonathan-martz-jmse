package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := os.Chdir("frontend")
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command("bun", "run", "generate")

	// Set the output to the terminal's standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	fmt.Println("Executing: bun run generate")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Command executed successfully!")

	err = os.Chdir("..")
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	// Define the local source directory and the remote destination
	sourceDir := "./frontend/.output/public/."
	remoteUser := "jonathp"
	remoteHost := "www466.your-server.de"
	remotePath := "public_html"
	remoteDest := fmt.Sprintf("%s@%s:%s", remoteUser, remoteHost, remotePath)

	// Build the SCP command
	cmd = exec.Command("scp", "-r", sourceDir, remoteDest)

	// Redirect standard output and error to the console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	fmt.Println("Running SCP command...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Files copied successfully!")
}
