package app

import (
	"testing"

	"github.com/henrioseptiano/carepedia-code-test/models"
)

func Test_Dequeue_EmptyPatient(t *testing.T) {
	queue.Dequeue()
}

func Test_Dequeue_WithPatient(t *testing.T) {
	patients := []models.Patient{{MRNumber: "MR2145", Gender: "F"}, {MRNumber: "MR2125", Gender: "M"}}
	//queue.Patients = patients
	for _, patient := range patients {
		queue.Enqueue(patient)
	}
	queue.Dequeue()
}

func Test_Dequeue_WithPatientAndRoundRobin(t *testing.T) {
	patients := []models.Patient{
		{MRNumber: "MR7724", Gender: "F"},
		{MRNumber: "MR3311", Gender: "F"},
		{MRNumber: "MR8214", Gender: "F"},
		{MRNumber: "MR9919", Gender: "M"},
		{MRNumber: "MR5241", Gender: "M"},
	}
	//queue.Patients = patients
	for _, patient := range patients {
		queue.Enqueue(patient)
	}
	queue.SetRoundRobin(true)
	queue.Dequeue()
}
