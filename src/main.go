package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("promptを入力してください: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("アプリケーションを終了します。")
			return
		}
		generatedText, err := handleOpenAi(text)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		handleSlack(generatedText)

	}
}
