package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/henrioseptiano/carepedia-code-test/app"
	"github.com/henrioseptiano/carepedia-code-test/models"
	"github.com/henrioseptiano/carepedia-code-test/utils"
)

func main() {
	var wg sync.WaitGroup
	commands := make(chan models.Command)
	done := make(chan bool)
	queue := app.NewQueue()
	wg.Add(1)
	go utils.QueueManager(commands, done, &wg, queue)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			close(done)
			wg.Wait()
			break
		}

		commandLine := scanner.Text()
		if strings.ToUpper(commandLine) == "EXIT" {
			close(done)
			wg.Wait()
			break
		}

		parts := strings.Split(commandLine, " ")
		commandType := parts[0]

		switch strings.ToUpper(commandType) {
		case "IN":
			if len(parts) != 3 {
				fmt.Println("Invalid input format. Please Use: IN <MRNumber> <Gender>")
				continue
			}
			mrNumber, gender := parts[1], parts[2]
			commands <- models.Command{Type: strings.ToUpper(commandType), Patient: &models.Patient{MRNumber: mrNumber, Gender: gender}}
			continue
		case "OUT", "ROUNDROBIN", "DEFAULT":
			commands <- models.Command{Type: strings.ToUpper(commandType)}
			continue
		default:
			fmt.Println("Unknown Command. Please Use one of [IN, OUT, ROUNDROBIN, DEFAULT, EXIT].")
			continue
		}
	}
}
