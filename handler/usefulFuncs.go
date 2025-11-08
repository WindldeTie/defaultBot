package handler

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func (h *Handler) console() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		command := strings.Split(text, " ")

		switch command[0] {
		case "exit":
			log.Println("Завершение работы бота...")
			os.Exit(0)
		default:
			log.Printf("Неизвестная команда: %s \n", text)
		}
	}
}
