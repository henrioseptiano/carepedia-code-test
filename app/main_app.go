package app

import "github.com/henrioseptiano/carepedia-code-test/models"

type Queue struct {
	Patients      []models.Patient
	RoundRobin    bool
	LastGenderOut string
}

func NewQueue() *Queue {
	return &Queue{
		Patients:   make([]models.Patient, 0),
		RoundRobin: false,
	}
}
