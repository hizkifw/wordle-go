package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hizkifw/wordle-go/wordle"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	wordle, err := wordle.NewWordle()
	if err != nil {
		panic(err)
	}

	fmt.Println("Make a guess")
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		res, err := wordle.SubmitGuess(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(string(res))
	}
}
