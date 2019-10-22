package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const red = 31
const green = 32

const backSpace = 127
const deleteWord = 23
const deleteAll = 21

// thank you @blinry
// https://stackoverflow.com/q/15159118
func setuptty() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func setOutputColor(color int) {
	fmt.Print("\u001b[" + strconv.Itoa(color) + "m")
}

func resetOutputColor() {
	setOutputColor(0)
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func deleteLastWord(s string) string {
	lastIndex := len(s) - 1
	for ; s[lastIndex] != ' '; lastIndex-- {
	}
	for ; s[lastIndex] == ' '; lastIndex-- {
	}
	return s[:lastIndex+1]
}

func main() {
	setuptty()

	b := make([]byte, 1)
	input := ""

	for {
		os.Stdin.Read(b)

		if b[0] == backSpace && len(input) > 0 {
			input = input[:len(input)-1]
		} else if b[0] == deleteWord {
			input = deleteLastWord(input)
		} else if b[0] == deleteAll {
			input = ""
		} else {
			input = input + string(b)
		}

		fmt.Println(input)
		clearScreen()

		parts := strings.Split(input, " ")
		command := parts[0]
		args := parts[1:]
		out, err := exec.Command(command, args...).Output()

		if err == nil {
			setOutputColor(green)
			fmt.Println(input + ": ")
			resetOutputColor()
			fmt.Print(string(out))
		} else {
			setOutputColor(red)
			fmt.Println(input + ": ")
			resetOutputColor()
			fmt.Print(string(out))
		}
	}
}
