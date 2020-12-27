package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var commitMode string
	var commitTag string
	var commitMessage string

	fmt.Println("Commit mode")
	fmt.Printf("\nfeat | fix | support | test | service | docs\n\n")
	fmt.Scanln(&commitMode)

	fmt.Println("Commit tag?(optional)")
	fmt.Printf("\nrefactor | style | important | small | etc.\n\n")
	fmt.Scanln(&commitTag)

	fmt.Println("Commit message?")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		commitMessage = scanner.Text()
	}
	fmt.Print(commitMode + " " + commitTag + " " + commitMessage + "|||\n")

	cmd := exec.Command("run", "commit", "--mode", commitMode, "--tag", commitTag, "--message", commitMessage)
	stdout, err := cmd.Output()
	fmt.Print(string(stdout))

	if err != nil {
		log.Fatal(err)
	}
}
