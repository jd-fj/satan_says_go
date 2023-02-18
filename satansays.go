package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	hello()

	for {
		// create new readline instance
		rl, err := readline.New("⦔ ")
		if err != nil {
			fmt.Println("read err!")
		}
		defer rl.Close()

		// read input
		text, err := rl.Readline()
		if err != nil {
			fmt.Println("28. ERR: ", err)
		}
		text = strings.TrimSpace(strings.ToLower(text))

		// exit command
		if text == "jesus is lord" {
			goodbye()
			break
		}

		switch text {
		default:
			_, filename, _, ok := runtime.Caller(0)
			if !ok {
				fmt.Println("Error getting audio resources")
				return
			}
			dir := filepath.Dir(filename)

			// Route working audio
			targetSayOutput := filepath.Join(dir, "tmp", "sayOutput.aiff")

			// make text into speech, put it in tmp folder
			sayOutputFile := exec.Command("say", "-r", "140", "-o", targetSayOutput, text)
			err = sayOutputFile.Run()
			if err != nil {
				fmt.Println("54. ERR: ", err)
			}

			targetSoxOutput := filepath.Join(dir, "tmp", "soxOutput.aiff")

			// lower pitch/speed, save to new file
			lowerOutputFile := exec.Command(
				"sox",
				targetSayOutput,
				targetSoxOutput,
				"speed",
				".75",
				"pitch",
				"-450",
			)
			_, err := lowerOutputFile.CombinedOutput()
			if err != nil {
				// log.Fatal("Fatal: failed to run sox command: \n", string(output))
				if strings.Contains(err.Error(), "exit status") {
					os.Exit(1)
				} else {
					log.Fatal("Failed to run sox command, have you brew installed sox?")
				}
			}

			// play the lowered audio file
			cmd2 := exec.Command("afplay", targetSoxOutput)
			err = cmd2.Run()
			if err != nil {
				fmt.Println("85. ERR: ", err)
			}
		}
	}
}

func hello() {
	fmt.Println("ⵅ    𝖍  𝖆  𝖏  𝖑    ⫛    𝖘  𝖆  𝖙  𝖆  𝖓    ⵅ")

	// Get the absolute path to the directory containing the source code file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting audio resources")
		return
	}
	dir := filepath.Dir(filename)

	// Get source audio files
	targetFilePath := filepath.Join(dir, "assets", "hello.mp3")

	cmd := exec.Command("afplay", targetFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("faild to play hello %v", string(output))
	}
}

func goodbye() {
	terminalWidth := readline.GetScreenWidth()
	for i := 0; i < terminalWidth/12; i++ {
		fmt.Print(" ♱ ")
	}
	fmt.Print("    𝔥 𝔢 𝔩 𝔩   ⫛   𝔞 𝔴 𝔞 𝔦 𝔱 𝔰    ")
	for i := 0; i < terminalWidth/12; i++ {
		fmt.Print(" ♱ ")
	}
	fmt.Println() // <- that's needed to not delete the line of crosses

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting audio resources")
		return
	}
	dir := filepath.Dir(filename)

	// Get source audio files
	targetFilePath := filepath.Join(dir, "assets", "goodbye.mp3")

	cmd := exec.Command("afplay", targetFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("faild to play hello %v", string(output))
	}
}