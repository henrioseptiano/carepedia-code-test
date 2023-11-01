package utils

import (
	"fmt"
	"sync"

	"github.com/henrioseptiano/carepedia-code-test/app"
	"github.com/henrioseptiano/carepedia-code-test/models"
)

func QueueManager(commands chan models.Command, done chan bool, wg *sync.WaitGroup, queue app.QueueInterface) {
	defer wg.Done()

	for {
		select {
		case command := <-commands:
			switch command.Type {
			case "IN":
				if err := queue.Enqueue(*command.Patient); err != nil {
					fmt.Println(err)
				}
				continue
			case "OUT":
				patient := queue.Dequeue()
				if patient == nil {
					fmt.Println("No Patients in Queue")
					continue
				}
				fmt.Printf("Send: %s %s\n", patient.MRNumber, patient.Gender)
			case "ROUNDROBIN":
				queue.SetRoundRobin(true)
			case "DEFAULT":
				queue.SetRoundRobin(false)
			}
		case <-done:
			return
		}
	}
}
