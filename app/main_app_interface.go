package app

import "github.com/henrioseptiano/carepedia-code-test/models"

type QueueInterface interface {
	Enqueue(p models.Patient) error
	DeQueue() *models.Patient
	SetRoundRobin(enable bool)
}
