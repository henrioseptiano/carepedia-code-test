package app

import (
	"fmt"

	"github.com/henrioseptiano/carepedia-code-test/models"
)

func (q *Queue) Enqueue(p models.Patient) error {
	for _, patient := range q.Patients {
		if patient.MRNumber == p.MRNumber {
			return fmt.Errorf("patient with " + p.MRNumber + " already in queue\n")
		}
	}
	q.Patients = append(q.Patients, p)
	return nil
}
