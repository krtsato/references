package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("START")

	start("り")
}

func start(word string) {
	fmt.Printf("最初の文字 : %s\n", word)
	text := input()
	judge(text, word)
}

func input() string {
	fmt.Println("文字を入力して下さい")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func judge(text, word string) {
	textRune := []rune(text)
	size := len(textRune)
	firstWord := string(textRune[0])
	lastWord := string(textRune[size-1])

	if firstWord != word {
		fmt.Printf("最初の文字が違います\n入力した最初の文字 : %s\nもう一度入力して下さい", firstWord)
		start(word)
	}

	if lastWord == "ん" {
		fmt.Println("最初の文字が「ん」です\n")
		end()
	}

	start(lastWord)
}

func end() {
	fmt.Println("END")
	os.Exit(0)
}
