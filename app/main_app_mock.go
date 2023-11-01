package app

import "github.com/henrioseptiano/carepedia-code-test/models"

type MockQueue struct {
	EnqueueCalled       bool
	DequeueCalled       bool
	SetRoundRobinCalled bool
	RoundRobinEnabled   bool
	LastGenderOut       string
	Patients            []models.Patient
}

func NewMockQueue(patients []models.Patient) *MockQueue {
	return &MockQueue{
		Patients: patients,
	}
}
