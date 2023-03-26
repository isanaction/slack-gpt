package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("promptを入力してください: ")
	text, _ := reader.ReadString('\n')
	handleOpenAi(text)
}
