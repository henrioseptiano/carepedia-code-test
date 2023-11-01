package app

import "github.com/henrioseptiano/carepedia-code-test/models"

type QueueInterface interface {
	Enqueue(p models.Patient) error
	Dequeue() *models.Patient
	SetRoundRobin(enable bool)
}
