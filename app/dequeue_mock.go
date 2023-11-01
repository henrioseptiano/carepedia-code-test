package app

import "github.com/henrioseptiano/carepedia-code-test/models"

func (m *MockQueue) Dequeue() *models.Patient {
	m.DequeueCalled = true

	if len(m.Patients) == 0 {
		return nil
	}

	if m.RoundRobinEnabled {
		for i, patient := range m.Patients {
			if m.LastGenderOut == "" || m.LastGenderOut != patient.Gender {
				m.LastGenderOut = patient.Gender
				m.Patients = append(m.Patients[:i], m.Patients[i+1:]...)
				return &patient
			}
		}
		m.LastGenderOut = ""
		return m.Dequeue()
	}

	patient := m.Patients[0]
	m.Patients = m.Patients[1:]
	return &patient
}
