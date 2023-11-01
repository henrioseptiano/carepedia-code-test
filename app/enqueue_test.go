package app

import (
	"testing"

	"github.com/henrioseptiano/carepedia-code-test/models"
)

func Test_Enqueue(t *testing.T) {

	patients := []models.Patient{{MRNumber: "MR2145", Gender: "F"}, {MRNumber: "MR2125", Gender: "M"}, {MRNumber: "MR2145", Gender: "M"}}
	//queue.Patients = patients
	for _, patient := range patients {
		queue.Enqueue(patient)
	}
}
