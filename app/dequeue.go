package app

import "github.com/henrioseptiano/carepedia-code-test/models"

func (q *Queue) DeQueue() *models.Patient {
	if len(q.Patients) == 0 {
		return nil
	}

	if q.RoundRobin {
		for i, patient := range q.Patients {
			if q.LastGenderOut == "" || q.LastGenderOut != patient.Gender {
				q.LastGenderOut = patient.Gender
				q.Patients = append(q.Patients[:i], q.Patients[i+1:]...)
				return &patient
			}
		}
	}

	patient := q.Patients[0]
	q.Patients = q.Patients[1:]
	return &patient
}
