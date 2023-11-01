package app

import "github.com/henrioseptiano/carepedia-code-test/models"

func (m *MockQueue) Enqueue(p models.Patient) error {
	m.EnqueueCalled = true
	m.Patients = append(m.Patients, p)
	return nil
}
