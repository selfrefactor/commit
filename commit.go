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
	fmt.Printf("\nfeat | fix | service | chore | test | docs\n\n")
	fmt.Scanln(&commitMode)

	fmt.Println("Commit tag?")
	fmt.Printf("\n🎏 refactor | 📦 dep | 🔥 important | 🆗 small | 🦠 issue\n")
	fmt.Printf("\n🎰 publish | 🏇 bump | 🐸 break | 🔪 deprecate | 💮 script\n")
	fmt.Printf("\n🥑 typings | 🎳 lint | 💋 style | 🐪 build | docs | changelog\n")
	fmt.Printf("\nexamples | method | prepublish | stop | usage\n")

	commitTagScanner := bufio.NewScanner(os.Stdin)

	if commitTagScanner.Scan() {
		commitTag = commitTagScanner.Text()
	}
	fmt.Println("Commit message?")

	commitMessageScanner := bufio.NewScanner(os.Stdin)
	if commitMessageScanner.Scan() {
		commitMessage = commitMessageScanner.Text()
	}
	fmt.Print(commitMode + " " + commitTag + " " + commitMessage + "|\n")
	return
	cmd := exec.Command("run", "commit", "--mode", commitMode, "--tag", commitTag, "--message", commitMessage)
	stdout, err := cmd.Output()
	fmt.Print(string(stdout))

	if err != nil {
		log.Fatal(err)
	}
}
