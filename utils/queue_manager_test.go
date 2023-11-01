package utils

import (
	"sync"
	"testing"

	"github.com/henrioseptiano/carepedia-code-test/app"
	"github.com/henrioseptiano/carepedia-code-test/models"
)

func TestQueueManagerRoundRobin(t *testing.T) {
	// Setup

	commands := make(chan models.Command)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	queue := app.NewMockQueue([]models.Patient{{MRNumber: "MR1001", Gender: "M"}, {MRNumber: "MR1002", Gender: "F"}})
	go QueueManager(commands, done, &wg, queue)

	// Test round-robin behavior
	commands <- models.Command{Type: "IN", Patient: &models.Patient{MRNumber: "MR2145", Gender: "F"}}
	commands <- models.Command{Type: "IN", Patient: &models.Patient{MRNumber: "MR5241", Gender: "M"}}
	commands <- models.Command{Type: "ROUNDROBIN"}
	commands <- models.Command{Type: "OUT"}
	commands <- models.Command{Type: "OUT"}
	//commands <- models.Command{Type: "DEFAULT"}
	// Signal to finish and wait
	close(done)
	wg.Wait()

	if !queue.RoundRobinEnabled {
		t.Errorf("Round robin was not enabled")
	}

	if !queue.DequeueCalled {
		t.Errorf("Dequeue was not called as expected")
	}
	if queue.LastGenderOut != "F" {
		t.Errorf("Expected last gender out to be 'F', got %s", queue.LastGenderOut)
	}
}

func TestQueueManagerRoundRobin_WithDefault(t *testing.T) {
	// Setup

	commands := make(chan models.Command)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	queue := app.NewMockQueue([]models.Patient{{MRNumber: "MR1001", Gender: "M"}, {MRNumber: "MR1002", Gender: "F"}})
	go QueueManager(commands, done, &wg, queue)

	// Test round-robin behavior
	commands <- models.Command{Type: "IN", Patient: &models.Patient{MRNumber: "MR2145", Gender: "F"}}
	commands <- models.Command{Type: "IN", Patient: &models.Patient{MRNumber: "MR5241", Gender: "M"}}
	commands <- models.Command{Type: "ROUNDROBIN"}
	commands <- models.Command{Type: "OUT"}
	commands <- models.Command{Type: "OUT"}
	commands <- models.Command{Type: "DEFAULT"}
	// Signal to finish and wait
	close(done)
	wg.Wait()

	if !queue.DequeueCalled {
		t.Errorf("Dequeue was not called as expected")
	}
}

func TestQueueManagerRoundRobin_WithEmptyDequeue(t *testing.T) {
	// Setup

	commands := make(chan models.Command)
	done := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(1)
	queue := app.NewMockQueue([]models.Patient{{MRNumber: "MR1001", Gender: "M"}, {MRNumber: "MR1002", Gender: "F"}})
	go QueueManager(commands, done, &wg, queue)

	// Test round-robin behavior
	commands <- models.Command{Type: "ROUNDROBIN"}
	commands <- models.Command{Type: "OUT"}
	commands <- models.Command{Type: "OUT"}
	commands <- models.Command{Type: "DEFAULT"}
	commands <- models.Command{Type: "OUT"}
	// Signal to finish and wait
	close(done)
	wg.Wait()

	if !queue.DequeueCalled {
		t.Errorf("Dequeue was not called as expected")
	}
}
