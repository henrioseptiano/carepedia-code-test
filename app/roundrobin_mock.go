package app

func (m *MockQueue) SetRoundRobin(enable bool) {
	m.SetRoundRobinCalled = true
	m.RoundRobinEnabled = enable
	m.LastGenderOut = ""
}
